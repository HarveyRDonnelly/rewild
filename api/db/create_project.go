package db

type CreateProjectDBRequest struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	PindropID         uuid_t `json:"pindrop_id"`
	TimelineID        uuid_t `json:"timeline_id"`
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	FollowerCount     int    `json:"follower_count"`
}

type CreateProjectDBResponse struct {
	ProjectID         uuid_t `json:"project_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	PindropID         uuid_t `json:"pindrop_id"`
	TimelineID        uuid_t `json:"timeline_id"`
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	FollowerCount     int    `json:"follower_count"`
}

func CreateProject(
	conn Connection,
	dbRequest CreateProjectDBRequest) CreateProjectDBResponse {

	var projectID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.projects (
					 name,
					 description, 
					 pindrop_id, 
					 timeline_id, 
					 discussion_board_id, 
					 follower_count
				 ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING project_id;`,
		dbRequest.Name,
		dbRequest.Description,
		dbRequest.PindropID,
		dbRequest.TimelineID,
		dbRequest.DiscussionBoardID,
		dbRequest.FollowerCount,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&projectID)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	dbResponse := CreateProjectDBResponse(
		GetProject(
			conn,
			GetProjectDBRequest{
				ProjectID: projectID,
			},
		),
	)

	return dbResponse

}
