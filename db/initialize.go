package db

import (
	"database/sql"

	db "github.com/zzihyeon/golang-server.git/db/sqlc"
)

var postgresqlClient *db.Queries

func Initialize(conn *sql.DB) {
	postgresqlClient = db.New(conn)
}

func GetClient() *db.Queries {
	return postgresqlClient
}
