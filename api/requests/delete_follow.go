package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

type DeleteFollowRequest struct {
	UserID uuid_t `json:"user_id"`
}

type DeleteFollowResponse struct {
	UserID uuid_t `json:"user_id"`
}

func deleteFollowRoute(r *gin.Engine) *gin.Engine {

	r.DELETE("/project/:project_id/follow", func(c *gin.Context) {

		var projectID = uuid.Must(uuid.Parse(c.Param("project_id")))
		var requestBody DeleteFollowRequest
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

		if doesFollow == true {
			db.DeleteFollow(
				DB,
				db.DeleteFollowDBRequest{
					ProjectID: projectID,
					UserID:    requestBody.UserID,
				},
			)
		}

		c.Status(http.StatusOK)
	})

	return r

}
