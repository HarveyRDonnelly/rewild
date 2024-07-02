package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

func getUserRoute(r *gin.Engine) *gin.Engine {
	r.GET("/user/:user_id/", func(c *gin.Context) {

		var userID = uuid.Must(uuid.Parse(c.Param("user_id")))
		var user = db.GetUser(DB, userID)

		c.JSON(
			http.StatusOK,
			user,
		)

	})

	return r
}
