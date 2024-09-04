package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type GetProjectResponse entities.Project

func getProjectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/project/:project_id/", func(c *gin.Context) {

		var projectID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("project_id"))),
			Valid: true,
		}

		if db.EntityExists(DB, projectID, "project") == false {
			error := entities.EntityNotExistsException("project")

			c.JSON(
				error.StatusCode,
				error,
			)
		} else {

			// Retrieve project info
			projectDBResponse := db.GetProject(
				DB,
				db.GetProjectDBRequest{
					ProjectID: projectID,
				},
			)

			newProject := db.ConstructProject(DB, projectDBResponse)

			c.JSON(
				http.StatusOK,
				GetProjectResponse(newProject),
			)

		}

	})

	return r
}
