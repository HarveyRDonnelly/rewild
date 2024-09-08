package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type UpdateProjectNamePostRequest struct {
	Name string `json:"name"`
}

type UpdateProjectNamePostResponse entities.Project

func updateProjectNameRoute(r *gin.Engine) *gin.Engine {

	r.PATCH("/project/:project_id/name", func(c *gin.Context) {

		var requestBody UpdateProjectNamePostRequest
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

		// Update project title
		updatedProjectDBResponse := db.UpdateProject(
			DB,
			db.UpdateProjectDBRequest{
				ProjectID:         projectID,
				Name:              requestBody.Name,
				Description:       projectDBResponse.Description,
				PindropID:         projectDBResponse.PindropID,
				TimelineID:        projectDBResponse.TimelineID,
				DiscussionBoardID: projectDBResponse.DiscussionBoardID,
				FollowerCount:     projectDBResponse.FollowerCount,
			},
		)

		// Construct updated project entity
		updatedProject := db.ConstructProject(
			DB,
			db.GetProjectDBResponse(updatedProjectDBResponse),
		)

		c.JSON(
			http.StatusOK,
			UpdateProjectNamePostResponse(updatedProject),
		)
	})

	return r

}
