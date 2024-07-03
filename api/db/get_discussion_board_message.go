package db

type GetDiscussionBoardMessageDBRequest struct {
	DiscussionBoardMessageID uuid_t `json:"discussion_board_message_id"`
}

type GetDiscussionBoardMessageDBResponse struct {
	DiscussionBoardMessageID uuid_t `json:"discussion_board_message_id"`
	ParentID                 uuid_t `json:"parent_id"`
	Body                     string `json:"body"`
}

func GetDiscussionBoardMessage(
	conn Connection,
	dbRequest GetDiscussionBoardMessageDBRequest) GetDiscussionBoardMessageDBResponse {

	var dbResponse GetDiscussionBoardMessageDBResponse
	dbResponse.DiscussionBoardMessageID = dbRequest.DiscussionBoardMessageID

	rows, err := conn.Gateway.Query(
		`SELECT parent_id, body FROM rewild.discussion_board_messages WHERE discussion_board_message_id=$1;`,
		nullIDString(dbRequest.DiscussionBoardMessageID))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&dbResponse.ParentID,
			&dbResponse.Body,
		)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return dbResponse

}
