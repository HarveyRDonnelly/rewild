package requests

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/option"
	"net/http"
	"rewild-it/api/db"
)

var app *firebase.App

type CreateUserSessionRequest struct {
	UserID   uuid_t `json:"user_id"`
	Password string `json:"password"`
}

type CreateUserSessionResponse struct {
	Token string `json:"token"`
}

func createUserSessionRoute(r *gin.Engine) *gin.Engine {
	r.POST("/login", func(c *gin.Context) {

		var requestBody CreateUserSessionRequest
		err := c.BindJSON(&requestBody)
		if err != nil {
			panic(err)
		}

		auth := db.GetAuth(
			DB,
			db.GetAuthDBRequest{
				UserID: requestBody.UserID,
			},
		)

		// Initialise firebase
		opt := option.WithCredentialsFile("./rewild-it-c744b-firebase-adminsdk-4pd7p-0f3e7e754d.json")
		app, err = firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			panic(err)
		}

		if checkPasswordHash(requestBody.Password, auth.Password) {
			client, err := app.Auth(context.Background())
			if err != nil {
				panic(err)
			}
			token, err := client.CustomToken(c, auth.UserID.UUID.String())
			if err != nil {
				panic(err)
			}
			c.JSON(
				http.StatusOK,
				CreateUserSessionResponse{
					Token: token,
				},
			)
		} else {
			c.Status(http.StatusUnauthorized)
		}


	})

	return r
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
