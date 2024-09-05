package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var UNPROTECTED_ROUTES = []string{"/login", "/user"}

var app *firebase.App

func init() {
	// Initialise firebase
	var err error
	// Load project absolute path
	var absolutePath, _ = os.LookupEnv("PROJECT_PATH")
	var firebaseOptsName, _ = os.LookupEnv("FIREBASE_OPTS_NAME")

	opt := option.WithCredentialsFile(absolutePath + "config/firebase/" + firebaseOptsName)
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
}

func AuthHandler(c *gin.Context) {

	idToken := c.Request.Header.Get("Authorization")
	if idToken != ""{
		idToken = strings.Split(idToken, "Bearer ")[1]
	}

	// Auth check
	client, err := app.Auth(c)
	if err != nil {
		panic(err)
	}
	routeIsProtected := true
	for _, route := range UNPROTECTED_ROUTES {
		if route == c.Request.URL.Path || route + "/" == c.Request.URL.Path{
			routeIsProtected = false
			break
		}
	}

	if routeIsProtected == true {
		_, err := client.VerifyIDToken(c, idToken)

		if err != nil {
			println(fmt.Sprintf("AUTH ERROR: %+v\n", err))
			c.Status(http.StatusForbidden)
			c.Abort()
		} else{
			c.Next()
		}
	} else {
		c.Next()
	}


}
