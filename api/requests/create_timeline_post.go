package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type CreateTimelinePostRequest struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Type     string `json:"type"`
	AuthorID uuid_t `json:"author_id"`
}

type CreateTimelinePostResponse entities.Timeline

func createTimelinePostRoute(r *gin.Engine) *gin.Engine {

	r.POST("/project/:project_id/timeline", func(c *gin.Context) {

		var requestBody CreateTimelinePostRequest
		var projectID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("project_id"))),
			Valid: true,
		}

		err := c.BindJSON(&requestBody)

		if err != nil {
			panic(err)
		}

		// Retrieve project
		projectDBResponse := db.GetProject(
			DB,
			db.GetProjectDBRequest{
				ProjectID: projectID,
			},
		)

		// Retrieve timeline
		timelineDBResponse := db.GetTimeline(
			DB,
			db.GetTimelineDBRequest{
				TimelineID: projectDBResponse.TimelineID,
			},
		)

		// Create new timeline post
		newTimelinePostDBResponse := db.CreateTimelinePost(
			DB,
			db.CreateTimelinePostDBRequest{
				NextID:   uuid.NullUUID{Valid: false},
				PrevID:   timelineDBResponse.TailID,
				Title:    requestBody.Title,
				Body:     requestBody.Body,
				AuthorID: requestBody.AuthorID,
				Type:     requestBody.Type,
			})

		// Retrieve timeline tail
		timelineTailPostDBResponse := db.GetTimelinePost(
			DB,
			db.GetTimelinePostDBRequest{
				TimelinePostID: timelineDBResponse.TailID,
			},
		)

		// Update timeline tail's next
		db.UpdateTimelinePost(
			DB,
			db.UpdateTimelinePostDBRequest{
				TimelinePostID: timelineTailPostDBResponse.TimelinePostID,
				NextID:         newTimelinePostDBResponse.TimelinePostID,
				PrevID:         timelineTailPostDBResponse.PrevID,
				Title:          timelineTailPostDBResponse.Title,
				Body:           timelineTailPostDBResponse.Body,
				AuthorID:       timelineTailPostDBResponse.AuthorID,
				Type:           timelineTailPostDBResponse.Type,
			},
		)

		// Update timeline tail
		updatedTimelineDBResponse := db.UpdateTimeline(
			DB,
			db.UpdateTimelineDBRequest{
				TimelineID: timelineDBResponse.TimelineID,
				HeadID:     timelineDBResponse.HeadID,
				TailID:     newTimelinePostDBResponse.TimelinePostID,
			},
		)

		// Construct updated timeline entity
		updatedTimeline := db.ConstructTimeline(
			DB,
			db.GetTimelineDBResponse(updatedTimelineDBResponse),
		)

		c.JSON(
			http.StatusCreated,
			CreateTimelinePostResponse(updatedTimeline),
		)
	})

	return r

}
