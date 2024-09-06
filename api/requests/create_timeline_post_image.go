package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

type CreateTimelinePostImageRequest struct {
	ImageID        uuid_t `json:"image_id"`
	TimelinePostID uuid_t `json:"timeline_post_id"`
}

type CreateTimelinePostImageResponse struct {
	ImageID        uuid_t `json:"image_id"`
	TimelinePostID uuid_t `json:"timeline_post_id"`
}

func createTimelinePostImageRoute(r *gin.Engine) *gin.Engine {

	r.POST("/timeline/post/:timeline_post_id/image/:image_id", func(c *gin.Context) {

		var requestBody CreateTimelinePostImageRequest

		requestBody.TimelinePostID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("timeline_post_id"))),
			Valid: true,
		}
		requestBody.ImageID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("image_id"))),
			Valid: true,
		}

		timelinePostImagesDBResponse := db.GetTimelinePostImages(
			DB,
			db.GetTimelinePostImagesDBRequest{
				TimelinePostID: requestBody.TimelinePostID,
			},
		)

		// TODO: Fix race condition here
		numImages := len(timelinePostImagesDBResponse.Images)

		db.CreateTimelinePostImage(
			DB,
			db.CreateTimelinePostImageDBRequest{
				ImageID:        requestBody.ImageID,
				ArrIndex:       numImages,
				TimelinePostID: requestBody.TimelinePostID,
			},
		)

		responseBody := CreateTimelinePostImageResponse(requestBody)

		c.JSON(
			http.StatusCreated,
			responseBody,
		)
	})

	return r

}
