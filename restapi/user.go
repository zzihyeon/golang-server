package restapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zzihyeon/golang-server.git/restapi/model"
)

func (server *Server) getMyInfo(ctx *gin.Context) {
	usrIdstr := ctx.Request.Header.Get("usrId")
	usrId, err := strconv.Atoi(usrIdstr)
	if err != nil || usrId < 0 {
		ctx.JSON(http.StatusOK, model.ErrorMessage{ErrCode: "1:Auth", Message: "로그인이 필요합니다."})
	} else {
		fmt.Println("get my info :  usrId", usrId)
		// usr, err := db.GetUser(ctx, int64(usrId))
		if err != nil {
			ctx.JSON(http.StatusOK, model.ErrorMessage{ErrCode: "2:NoUser", Message: "사용자 정보가 없습니다. !!"})
		} else {
			ctx.JSON(http.StatusOK, nil)
		}
	}
}
