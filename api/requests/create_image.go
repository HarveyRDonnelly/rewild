package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type CreateImageRequest struct {
	AltText string `json:"alt_text"`
}

type CreateImageResponse entities.Image

func createImageRoute(r *gin.Engine) *gin.Engine {

	r.POST("/timeline/post/:timeline_post_id/image", func(c *gin.Context) {

		var timelinePostID = uuid.Must(uuid.Parse(c.Param("timeline_post_id")))

		imageFile, _ := c.FormFile("image")
		altText, _ := c.GetPostForm("alt_text")

		dbResponse := db.CreateImage(
			DB,
			db.CreateImageDBRequest{
				TimelinePostID: timelinePostID,
				AltText:        altText,
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
			newImage,
		)
	})

	return r

}
