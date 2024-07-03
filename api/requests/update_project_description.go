package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type UpdateProjectDescriptionPostRequest struct {
	Description string `json:"description"`
}

type UpdateProjectDescriptionPostResponse entities.Project

func updateProjectDescriptionRoute(r *gin.Engine) *gin.Engine {

	r.PATCH("/project/:project_id/description", func(c *gin.Context) {

		var requestBody UpdateProjectDescriptionPostRequest
		var projectID = uuid.Must(uuid.Parse(c.Param("project_id")))

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
				Name:              projectDBResponse.Name,
				Description:       requestBody.Description,
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
