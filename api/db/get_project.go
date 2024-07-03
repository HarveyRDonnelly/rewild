package db

type GetProjectDBRequest struct {
	ProjectID uuid_t `json:"project_id"`
}

type GetProjectDBResponse struct {
	ProjectID         uuid_t `json:"project_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	PindropID         uuid_t `json:"pindrop_id"`
	TimelineID        uuid_t `json:"timeline_id"`
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	FollowerCount     int    `json:"follower_count"`
}

func GetProject(conn Connection, dbRequest GetProjectDBRequest) GetProjectDBResponse {

	var dbResponse GetProjectDBResponse
	dbResponse.ProjectID = dbRequest.ProjectID

	rows, err := conn.Gateway.Query(
		`SELECT name, description, pindrop_id, timeline_id, discussion_board_id, follower_count 
				FROM rewild.projects WHERE project_id=$1;`,
		nullIDString(dbRequest.ProjectID))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&dbResponse.Name,
			&dbResponse.Description,
			&dbResponse.PindropID,
			&dbResponse.TimelineID,
			&dbResponse.DiscussionBoardID,
			&dbResponse.FollowerCount)

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
