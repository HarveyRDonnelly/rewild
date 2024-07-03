package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

type CreateProjectRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	PindropLatitude  float64  `json:"pindrop_latitude"`
	PindropLongitude float64  `json:"pindrop_longitude"`
	Followers        []uuid_t `json:"followers"`
}

type CreateProjectResponse struct {
	ProjectID         uuid_t `json:"project_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	PindropID         uuid_t `json:"pindrop_id"`
	TimelineID        uuid_t `json:"timeline_id"`
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	FollowerCount     int    `json:"follower_count"`
}

func createProjectRoute(r *gin.Engine) *gin.Engine {

	r.POST("/project/", func(c *gin.Context) {

		var requestBody CreateProjectRequest

		err := c.BindJSON(&requestBody)

		if err != nil {
			panic(err)
		}

		// Create discussion board message
		discussionBoardMessageDBResponse := db.CreateDiscussionBoardMessage(
			DB,
			db.CreateDiscussionBoardMessageDBRequest{
				ParentID: uuid.Nil,
				Body:     "",
			},
		)

		// Create discussion board
		discussionBoardDBResponse := db.CreateDiscussionBoard(
			DB,
			db.CreateDiscussionBoardDBRequest{
				RootID: discussionBoardMessageDBResponse.DiscussionBoardMessageID,
			},
		)

		// Create initial timeline post
		timelinePostDBResponse := db.CreateTimelinePost(
			DB,
			db.CreateTimelinePostDBRequest{
				NextID: uuid.Nil,
				PrevID: uuid.Nil,
				Title:  "Let's start rewilding!",
				Body:   "",
			})

		// Create timeline
		timelineDBResponse := db.CreateTimeline(
			DB,
			db.CreateTimelineDBRequest{
				HeadID: timelinePostDBResponse.TimelinePostID,
				TailID: timelinePostDBResponse.TimelinePostID,
			},
		)

		// Create pindrop
		pindropDBResponse := db.CreatePindrop(
			DB,
			db.CreatePindropDBRequest{
				Latitude:  requestBody.PindropLatitude,
				Longitude: requestBody.PindropLongitude,
			},
		)

		// Create project
		projectDBResponse := db.CreateProject(
			DB,
			db.CreateProjectDBRequest{
				Name:              requestBody.Name,
				Description:       requestBody.Description,
				PindropID:         pindropDBResponse.PindropID,
				TimelineID:        timelineDBResponse.TimelineID,
				DiscussionBoardID: discussionBoardDBResponse.DiscussionBoardID,
				FollowerCount:     len(requestBody.Followers),
			},
		)

		// Add follower relationships
		for _, followerID := range requestBody.Followers {
			db.CreateFollow(
				DB,
				db.CreateFollowDBRequest{
					UserID:    followerID,
					ProjectID: projectDBResponse.ProjectID,
				},
			)
		}

		// Create response object
		newProject := CreateProjectResponse{
			ProjectID:     projectDBResponse.ProjectID,
			Name:          projectDBResponse.Name,
			Description:   projectDBResponse.Description,
			PindropID:     projectDBResponse.PindropID,
			TimelineID:    projectDBResponse.TimelineID,
			FollowerCount: projectDBResponse.FollowerCount,
		}

		c.JSON(
			http.StatusCreated,
			newProject,
		)
	})

	return r

}
