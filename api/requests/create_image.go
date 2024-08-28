package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type CreateImageResponse entities.Image

func createImageRoute(r *gin.Engine) *gin.Engine {

	r.POST("/image/", func(c *gin.Context) {

		imageFile, _ := c.FormFile("image")
		altText, _ := c.GetPostForm("alt_text")

		dbResponse := db.CreateImage(
			DB,
			db.CreateImageDBRequest{
				AltText: altText,
			},
		)

		err := c.SaveUploadedFile(imageFile, "./res/"+dbResponse.ImageID.String()+".png")

		if err != nil {
			panic(err)
		}

		newImage := db.ConstructImage(
			DB,
			db.GetImageDBResponse(dbResponse))

		c.JSON(
			http.StatusCreated,
			CreateImageResponse(newImage),
		)
	})

	return r

}
