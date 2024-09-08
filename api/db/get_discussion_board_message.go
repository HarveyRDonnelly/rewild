package db

type GetDiscussionBoardMessageDBRequest struct {
	DiscussionBoardMessageID uuid_t `json:"discussion_board_message_id"`
}

type GetDiscussionBoardMessageDBResponse struct {
	DiscussionBoardMessageID uuid_t `json:"discussion_board_message_id"`
	ParentID                 uuid_t `json:"parent_id"`
	Body                     string `json:"body"`
	AuthorID                 uuid_t `json:"author_id"`
	CreatedTS                string `json:"created_ts"`
}

func GetDiscussionBoardMessage(
	conn Connection,
	dbRequest GetDiscussionBoardMessageDBRequest) GetDiscussionBoardMessageDBResponse {

	var dbResponse GetDiscussionBoardMessageDBResponse
	dbResponse.DiscussionBoardMessageID = dbRequest.DiscussionBoardMessageID

	rows, err := conn.Gateway.Query(
		`SELECT parent_id, body, author_id, created_ts FROM rewild.discussion_board_messages WHERE discussion_board_message_id=$1;`,
		dbRequest.DiscussionBoardMessageID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&dbResponse.ParentID,
			&dbResponse.Body,
			&dbResponse.AuthorID,
			&dbResponse.CreatedTS,
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
