package requests

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/option"
	"net/http"
	"os"
	"rewild-it/api/db"
)

var app *firebase.App

type CreateUserSessionRequest struct {
	Username   string `json:"username"`
	Password string `json:"password"`
}

type CreateUserSessionResponse struct {
	Token string `json:"token"`
}

func createUserSessionRoute(r *gin.Engine) *gin.Engine {
	r.POST("/login", func(c *gin.Context) {

		// Load project absolute path
		var absolutePath, _ = os.LookupEnv("PROJECT_PATH")
		var firebaseOptsName, _ = os.LookupEnv("FIREBASE_OPTS_NAME")

		var requestBody CreateUserSessionRequest
		err := c.BindJSON(&requestBody)
		if err != nil {
			panic(err)
		}

		userID := db.FindUserIDByUsername(
			DB,
			requestBody.Username)

		auth := db.GetAuth(
			DB,
			db.GetAuthDBRequest{
				UserID: userID,
			},
		)

		// Initialise firebase
		opt := option.WithCredentialsFile(absolutePath + "config/firebase/" + firebaseOptsName)
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
