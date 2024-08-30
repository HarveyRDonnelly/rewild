package db

type DeleteFollowDBRequest struct {
	UserID    uuid_t `json:"user_id"`
	ProjectID uuid_t `json:"project_id"`
}

type DeleteFollowDBResponse struct {
	UserID    uuid_t `json:"user_id"`
	ProjectID uuid_t `json:"project_id"`
}

func DeleteFollow(
	conn Connection,
	dbRequest DeleteFollowDBRequest) DeleteFollowDBResponse {

	rows, err := conn.Gateway.Query(
		`DELETE FROM rewild.follows
				WHERE user_id=$1 AND project_id=$2;`,
		nullIDString(dbRequest.UserID),
		nullIDString(dbRequest.ProjectID),
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	dbResponse := DeleteFollowDBResponse(dbRequest)

	return dbResponse

}
