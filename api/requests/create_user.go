package requests

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rewild-it/api/db"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type CreateUserResponse struct {
	UserID    uuid_t   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Username  string   `json:"username"`
	Follows   []uuid_t `json:"follows"`
}

func createUserRoute(r *gin.Engine) *gin.Engine {

	r.POST("/user/", func(c *gin.Context) {

		var requestBody CreateUserRequest
		err := c.BindJSON(&requestBody)
		if err != nil {
			panic(err)
		}

		dbResponse := db.CreateUser(
			DB,
			db.CreateUserDBRequest{
				FirstName: requestBody.FirstName,
				LastName:  requestBody.LastName,
				Email:     requestBody.Email,
				Username:  requestBody.Username,
			},
		)

		db.CreateAuth(
			DB,
			db.CreateAuthDBRequest{
				UserID: dbResponse.UserID,
				Password: hashPassword(requestBody.Password),
			},
		)

		newUser := CreateUserResponse{
			UserID:    dbResponse.UserID,
			FirstName: dbResponse.FirstName,
			LastName:  dbResponse.LastName,
			Email:     dbResponse.Email,
			Username:  dbResponse.Username,
			Follows:   dbResponse.Follows,
		}

		c.JSON(
			http.StatusCreated,
			newUser,
		)
	})

	return r

}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
