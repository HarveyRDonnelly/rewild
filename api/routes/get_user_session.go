package routes

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"net/http"
	"os"
)

type GetUserSessionResponse struct {
	Token string `json:"token"`
}

func getUserSessionRoute(r *gin.Engine) *gin.Engine {
	r.GET("/user/:user_id/token", func(c *gin.Context) {

		var userID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("user_id"))),
			Valid: true,
		}

		// Load project absolute path
		var absolutePath, _ = os.LookupEnv("PROJECT_PATH")
		var firebaseOptsName, _ = os.LookupEnv("FIREBASE_OPTS_NAME")

		// Initialise firebase
		opt := option.WithCredentialsFile(absolutePath + "config/firebase/" + firebaseOptsName)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			panic(err)
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			panic(err)
		}
		token, err := client.CustomToken(c, userID.UUID.String())
		if err != nil {
			panic(err)
		}
		c.JSON(
			http.StatusOK,
			GetUserSessionResponse{
				Token: token,
			},
		)

	})

	return r
}
