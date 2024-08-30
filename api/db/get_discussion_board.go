package db

type GetDiscussionBoardDBRequest struct {
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
}

type GetDiscussionBoardDBResponse struct {
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	RootID            uuid_t `json:"root_id"`
}

func GetDiscussionBoard(
	conn Connection,
	dbRequest GetDiscussionBoardDBRequest) GetDiscussionBoardDBResponse {

	var dbResponse GetDiscussionBoardDBResponse
	dbResponse.DiscussionBoardID = dbRequest.DiscussionBoardID

	rows, err := conn.Gateway.Query(
		`SELECT root_id FROM rewild.discussion_boards WHERE discussion_board_id=$1;`,
		dbRequest.DiscussionBoardID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&dbResponse.RootID,
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
