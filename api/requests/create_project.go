package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rewild-it/api/db"
)

type CreateProjectRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	PindropLatitude  uuid_t   `json:"pindrop_latitude"`
	PindropLongitude uuid_t   `json:"pindrop_longitude"`
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

		dbResponse := db.CreateProject(
			DB,
			db.CreateProjectDBRequest{
				FirstName: requestBody.FirstName,
				LastName:  requestBody.LastName,
				Email:     requestBody.Email,
				Username:  requestBody.Username,
			},
		)

		newUser := CreateUserResponse{
			UserID:    dbResponse.UserID,
			FirstName: dbResponse.FirstName,
			LastName:  dbResponse.LastName,
			Email:     dbResponse.Email,
			Username:  dbResponse.Username,
			Follows:   dbResponse.Follows,
		}

		c.JSON(
			http.StatusCreated,
			newUser,
		)
	})

	return r

}
