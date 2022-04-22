package restapi

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/o1egl/paseto"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}

	router := gin.Default()

	router.Use(AuthWare)

	// Test apis
	router.GET("/test/token", server.getMyInfo)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func AuthWare(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	// token, err := ctx.Cookie("Authorization")
	fmt.Println("MWARE :: AuthWare : token", token)
	if !strings.Contains(token, "Bearer") {
		fmt.Println("MWARE :: AuthWare  :  No Token")
		ctx.Next()
		return
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	usrId := VerifyPasetoToken(token)
	fmt.Println("MWARE :: AuthWare : usr id", usrId)
	ctx.Request.Header.Set("auth", "true")
	ctx.Request.Header.Set("usrId", strconv.Itoa(usrId))
	ctx.Next()
}

func VerifyPasetoToken(token string) int {
	var newJsonToken paseto.JSONToken
	var newFooter string
	v2 := paseto.NewV2()
	pub, _ := hex.DecodeString("1212312")
	publicKey := ed25519.PublicKey(pub)

	err := v2.Verify(token, publicKey, &newJsonToken, &newFooter)
	if err != nil {
		return -1
	}
	uid, err := strconv.Atoi(newJsonToken.Get("uid"))
	if err != nil {
		return -2
	}
	return uid
}
