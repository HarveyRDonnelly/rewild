package api

import (
	"github.com/gin-gonic/gin"
	"rewild-it/api/db"
	"rewild-it/api/routes"
)

var r *gin.Engine

func SetDB(db db.Connection) {
	routes.SetDB(db)
}

func init() {
	r = routes.Create()
}

func Run() {
	r.Run()
}
