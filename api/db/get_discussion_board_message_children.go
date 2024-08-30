package db

type GetDiscussionBoardMessageChildrenDBRequest struct {
	ParentMessageID uuid_t `json:"parent_message_id"`
}

type GetDiscussionBoardMessageChildrenDBResponse struct {
	ChildMessages []GetDiscussionBoardMessageDBResponse `json:"child_messages"`
}

func GetDiscussionBoardMessageChildren(
	conn Connection,
	dbRequest GetDiscussionBoardMessageChildrenDBRequest) GetDiscussionBoardMessageChildrenDBResponse {

	var dbResponse GetDiscussionBoardMessageChildrenDBResponse
	var currDiscussionBoardMessageDBResponse GetDiscussionBoardMessageDBResponse

	currDiscussionBoardMessageDBResponse.ParentID = dbRequest.ParentMessageID

	rows, err := conn.Gateway.Query(
		`SELECT discussion_board_message_id, body, author_id FROM rewild.discussion_board_messages WHERE parent_id=$1;`,
		nullIDString(dbRequest.ParentMessageID))

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&currDiscussionBoardMessageDBResponse.DiscussionBoardMessageID,
			&currDiscussionBoardMessageDBResponse.Body,
			&currDiscussionBoardMessageDBResponse.AuthorID,
		)

		dbResponse.ChildMessages = append(dbResponse.ChildMessages, currDiscussionBoardMessageDBResponse)

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
