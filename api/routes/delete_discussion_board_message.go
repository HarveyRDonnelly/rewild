package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

func deleteDiscussionBoardMessageRoute(r *gin.Engine) *gin.Engine {

	r.DELETE("/discussion/message/:discussion_board_message_id", func(c *gin.Context) {

		var discussionBoardMessageID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("discussion_board_message_id"))),
			Valid: true,
		}

		db.DeleteDiscussionBoardMessage(
			DB,
			db.DeleteDiscussionBoardMessageDBRequest{
				DiscussionBoardMessageID: discussionBoardMessageID,
			},
		)

		c.Status(http.StatusOK)
	})

	return r

}
