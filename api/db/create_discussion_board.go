package db

type CreateDiscussionBoardDBRequest struct {
	RootID uuid_t `json:"root_id"`
}

type CreateDiscussionBoardDBResponse struct {
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	RootID            uuid_t `json:"root_id"`
}

func CreateDiscussionBoard(
	conn Connection,
	dbRequest CreateDiscussionBoardDBRequest) CreateDiscussionBoardDBResponse {

	var discussionBoardID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.discussion_boards (root_id)
				VALUES ($1) RETURNING discussion_board_id;`,
		nullIDString(dbRequest.RootID),
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&discussionBoardID)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	dbResponse := CreateDiscussionBoardDBResponse(
		GetDiscussionBoard(
			conn,
			GetDiscussionBoardDBRequest{
				DiscussionBoardID: discussionBoardID,
			},
		),
	)

	return dbResponse

}
