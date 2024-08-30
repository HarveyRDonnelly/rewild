package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

type CreateFollowRequest struct {
	UserID uuid_t `json:"user_id"`
}

type CreateFollowResponse struct {
	UserID uuid_t `json:"user_id"`
}

func createFollowRoute(r *gin.Engine) *gin.Engine {

	r.POST("/project/:project_id/follow", func(c *gin.Context) {

		var projectID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("project_id"))),
			Valid: true,
		}
		var requestBody CreateFollowRequest
		err := c.BindJSON(&requestBody)
		if err != nil {
			panic(err)
		}

		userDBResponse := db.GetUser(
			DB,
			db.GetUserDBRequest{
				UserID: requestBody.UserID,
			},
		)

		doesFollow := false
		for i := 0; i < len(userDBResponse.Follows); i++ {
			if userDBResponse.Follows[i] == projectID {
				doesFollow = true
				break
			}
		}

		if doesFollow == false {
			db.CreateFollow(
				DB,
				db.CreateFollowDBRequest{
					ProjectID: projectID,
					UserID:    requestBody.UserID,
				},
			)

			projectDBResponse := db.GetProject(
				DB,
				db.GetProjectDBRequest{
					ProjectID: projectID,
				},
			)

			projectDBResponse.FollowerCount += 1

			newProjectDBResponse := db.UpdateProject(
				DB,
				db.UpdateProjectDBRequest(projectDBResponse),
			)

			c.JSON(
				http.StatusCreated,
				db.ConstructProject(
					DB,
					db.GetProjectDBResponse(newProjectDBResponse),
				),
			)
		} else {
			c.JSON(
				http.StatusOK,
				db.ConstructProject(
					DB,
					db.GetProject(
						DB,
						db.GetProjectDBRequest{
							ProjectID: projectID,
						},
					),
				),
			)
		}
	})

	return r

}
