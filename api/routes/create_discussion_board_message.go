package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rewild-it/api/db"
	"rewild-it/api/entities"
)

type CreateDiscussionBoardMessageRequest struct {
	ParentID uuid_t `json:"parent_id"`
	Body     string `json:"body"`
	AuthorID uuid_t `json:"author_id"`
}

type CreateDiscussionBoardMessageResponse entities.DiscussionBoardMessage

func createDiscussionBoardMessageRoute(r *gin.Engine) *gin.Engine {

	r.POST("/discussion", func(c *gin.Context) {

		var requestBody CreateDiscussionBoardMessageRequest

		err := c.BindJSON(&requestBody)
		if err != nil {
			panic(err)
		}

		dbResponse := db.CreateDiscussionBoardMessage(
			DB,
			db.CreateDiscussionBoardMessageDBRequest(requestBody),
		)

		newMessage := db.ConstructDiscussionBoardMessageLimited(
			DB,
			db.GetDiscussionBoardMessageDBResponse(dbResponse),
			-1,
			1,
		)

		c.JSON(
			http.StatusCreated,
			CreateDiscussionBoardMessageResponse(newMessage),
		)
	})

	return r

}
