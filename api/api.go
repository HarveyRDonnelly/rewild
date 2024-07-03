package api

import (
	"github.com/gin-gonic/gin"
	"rewild-it/api/db"
	"rewild-it/api/requests"
)

var r *gin.Engine

func SetDB(db db.Connection) {
	requests.SetDB(db)
}

func init() {
	r = requests.Create()
}

func Run() {
	r.Run()
}
