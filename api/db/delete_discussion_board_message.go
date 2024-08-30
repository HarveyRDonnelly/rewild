package db

type DeleteDiscussionBoardMessageDBRequest struct {
	DiscussionBoardMessageID uuid_t `json:"discussion_board_message_id"`
}

type DeleteDiscussionBoardMessageDBResponse struct {
	DiscussionBoardMessageID uuid_t `json:"discussion_board_message_id"`
}

func DeleteDiscussionBoardMessage(
	conn Connection,
	dbRequest DeleteDiscussionBoardMessageDBRequest) DeleteDiscussionBoardMessageDBResponse {

	var dbResponse DeleteDiscussionBoardMessageDBResponse
	dbResponse.DiscussionBoardMessageID = dbRequest.DiscussionBoardMessageID

	rows, err := conn.Gateway.Query(
		`DELETE FROM rewild.discussion_board_messages 
                                  WHERE discussion_board_message_id=$1;`,
		nullIDString(dbRequest.DiscussionBoardMessageID))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	dbResponse.DiscussionBoardMessageID = dbRequest.DiscussionBoardMessageID

	return dbResponse

}
