package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type UpdateImageAltTextRequest struct {
	AltText string `json:"alt_text"`
}

type UpdateImageAltTextResponse entities.Image

func updateImageAltTextRoute(r *gin.Engine) *gin.Engine {

	r.PATCH("/image/:image_id", func(c *gin.Context) {

		var requestBody UpdateImageAltTextRequest
		var imageID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("image_id"))),
			Valid: true,
		}

		err := c.BindJSON(&requestBody)

		if err != nil {
			panic(err)
		}

		// Update project title
		updatedImageDBResponse := db.UpdateImage(
			DB,
			db.UpdateImageDBRequest{
				ImageID: imageID,
				AltText: requestBody.AltText,
			},
		)

		// Construct updated project entity
		updatedImage := db.ConstructImage(
			DB,
			db.GetImageDBResponse(updatedImageDBResponse),
		)

		c.JSON(
			http.StatusOK,
			UpdateImageAltTextResponse(updatedImage),
		)
	})

	return r

}
