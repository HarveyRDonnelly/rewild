package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var app *firebase.App

func init() {
	// Initialise firebase
	var err error
	opt := option.WithCredentialsFile("./rewild-it-c744b-firebase-adminsdk-4pd7p-0f3e7e754d.json")
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		print("POOP")
	}
}

func AuthHandler(c *gin.Context) {

	print("< AUTH MIDDLEWARE PLACEHOLDER >\n")
	idToken := c.Request.Header.Get("Authorization")

	// Auth check
	client, err := app.Auth(c)
	if err != nil {
		print(fmt.Sprintf("error getting Auth client: %v\n", err))
	}

	token, err := client.VerifyIDToken(c, idToken)
	if err != nil {
		print(fmt.Sprintf("error verifying ID token: %v\n", err))
	}

	print(fmt.Sprintf("Verified ID token: %v\n", token))

	c.Next()
}
