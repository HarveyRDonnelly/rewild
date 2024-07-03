package db

type CreateFollowDBRequest struct {
	UserID    uuid_t `json:"user_id"`
	ProjectID uuid_t `json:"project_id"`
}

type CreateFollowDBResponse struct {
	UserID    uuid_t `json:"user_id"`
	ProjectID uuid_t `json:"project_id"`
}

func CreateFollow(
	conn Connection,
	dbRequest CreateFollowDBRequest) CreateFollowDBResponse {

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.follows (user_id, project_id)
				VALUES ($1, $2);`,
		dbRequest.UserID,
		dbRequest.ProjectID,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	dbResponse := CreateFollowDBResponse(dbRequest)

	return dbResponse

}
