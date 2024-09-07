package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

type GetUserResponse struct {
	UserID    uuid_t   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Follows   []uuid_t `json:"follows"`
}

func getUserRoute(r *gin.Engine) *gin.Engine {
	r.GET("/user/:user_id/", func(c *gin.Context) {

		var userID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("user_id"))),
			Valid: true,
		}
		var user = db.GetUser(
			DB,
			db.GetUserDBRequest{
				UserID: userID,
			},
		)

		c.JSON(
			http.StatusOK,
			GetUserResponse(user),
		)

	})

	return r
}
