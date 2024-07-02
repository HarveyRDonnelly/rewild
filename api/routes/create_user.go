package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/requests"
)

func createUserRoute(r *gin.Engine) *gin.Engine {

	r.POST("/user/", func(c *gin.Context) {

		var requestBody requests.CreateUserRequest

		err := c.BindJSON(&requestBody)

		if err != nil {
			panic(err)
		}

		var firstName = requestBody.FirstName
		var lastName = requestBody.LastName
		var username = requestBody.Username
		var email = requestBody.Email

		db.CreateUser(
			DB,
			firstName,
			lastName,
			username,
			email)

		c.Status(http.StatusCreated)
	})

	return r

}
