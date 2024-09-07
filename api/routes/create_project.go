package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type CreateProjectRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	PindropLatitude  float64  `json:"pindrop_latitude"`
	PindropLongitude float64  `json:"pindrop_longitude"`
	Followers        []uuid_t `json:"followers"`
}

type CreateProjectResponse entities.Project

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
				ParentID: uuid.NullUUID{Valid: false},
				AuthorID: uuid.NullUUID{Valid: false},
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
				NextID:   uuid.NullUUID{Valid: false},
				PrevID:   uuid.NullUUID{Valid: false},
				Title:    "Let's start rewilding!",
				Body:     "",
				Type:     "initial",
				AuthorID: uuid.NullUUID{Valid: false},
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

		newProjectDBResponse := db.GetProject(
			DB,
			db.GetProjectDBRequest{
				ProjectID: projectDBResponse.ProjectID,
			})

		c.JSON(
			http.StatusCreated,
			CreateProjectResponse(db.ConstructProject(
				DB,
				newProjectDBResponse)),
		)
	})

	return r

}
