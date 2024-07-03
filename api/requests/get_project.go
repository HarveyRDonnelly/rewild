package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

type GetProjectResponse struct {
	ProjectID         uuid_t `json:"project_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	PindropID         uuid_t `json:"pindrop_id"`
	TimelineID        uuid_t `json:"timeline_id"`
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	FollowerCount     int    `json:"follower_count"`
}

func getProjectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/project/:project_id/", func(c *gin.Context) {

		var projectID = uuid.Must(uuid.Parse(c.Param("project_id")))
		var project = db.GetProject(
			DB,
			db.GetProjectDBRequest{
				ProjectID: projectID,
			},
		)

		c.JSON(
			http.StatusOK,
			GetProjectResponse(project),
		)

	})

	return r
}
