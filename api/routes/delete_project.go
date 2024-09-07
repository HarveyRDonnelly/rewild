package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

func deleteProjectRoute(r *gin.Engine) *gin.Engine {

	r.DELETE("/project/:project_id/", func(c *gin.Context) {

		var projectID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("project_id"))),
			Valid: true,
		}

		project := db.ConstructProject(
			DB,
			db.GetProject(
				DB,
				db.GetProjectDBRequest{
					ProjectID: projectID,
				},
			),
		)

		db.DeleteFollows(
			DB,
			db.DeleteFollowsDBRequest{
				ProjectID: projectID,
			},
		)

		db.DeleteProject(
			DB,
			db.DeleteProjectDBRequest{
				ProjectID: projectID,
			},
		)

		db.DeletePindrop(
			DB,
			db.DeletePindropDBRequest{
				PindropID: project.Pindrop.PindropID,
			},
		)

		c.Status(http.StatusOK)
	})

	return r

}
