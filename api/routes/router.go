package routes

import (
	"github.com/gin-gonic/gin"
	"rewild-it/api/db"
)

var DB db.Connection

func SetDB(db db.Connection) {
	DB = db
}

func Create() *gin.Engine {
	r := gin.Default()

	r = getUserRoute(r)
	r = createUserRoute(r)

	return r
}
