package db

type CreateDiscussionBoardMessageDBRequest struct {
	ParentID uuid_t `json:"parent_id"`
	Body     string `json:"body"`
}

type CreateDiscussionBoardMessageDBResponse struct {
	DiscussionBoardMessageID uuid_t `json:"discussion_board_message_id"`
	ParentID                 uuid_t `json:"parent_id"`
	Body                     string `json:"body"`
}

func CreateDiscussionBoardMessage(
	conn Connection,
	dbRequest CreateDiscussionBoardMessageDBRequest) CreateDiscussionBoardMessageDBResponse {

	var discussionBoardMessageID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.discussion_board_messages (parent_id, body)
				VALUES ($1, $2) RETURNING discussion_board_message_id;`,
		nullIDString(dbRequest.ParentID),
		dbRequest.Body,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&discussionBoardMessageID)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	dbResponse := CreateDiscussionBoardMessageDBResponse(
		GetDiscussionBoardMessage(
			conn,
			GetDiscussionBoardMessageDBRequest{
				DiscussionBoardMessageID: discussionBoardMessageID,
			},
		),
	)

	return dbResponse

}
