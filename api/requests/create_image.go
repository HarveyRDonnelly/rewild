package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type CreateImageResponse entities.Image

func createImageRoute(r *gin.Engine) *gin.Engine {

	// Load project absolute path
	var absolutePath, _ = os.LookupEnv("PROJECT_PATH")

	r.POST("/image/", func(c *gin.Context) {

		imageFile, err := c.FormFile("image_file")
		if err != nil {
			panic(err)
		}

		altText, _ := c.GetPostForm("alt_text")

		dbResponse := db.CreateImage(
			DB,
			db.CreateImageDBRequest{
				AltText: altText,
			},
		)

		err = c.SaveUploadedFile(imageFile, absolutePath+"res/"+dbResponse.ImageID.UUID.String()+".png")

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
