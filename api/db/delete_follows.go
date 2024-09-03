package db

type DeleteFollowsDBRequest struct {
	ProjectID uuid_t `json:"project_id"`
}

type DeleteFollowsDBResponse struct {
	ProjectID uuid_t `json:"project_id"`
}

func DeleteFollows(
	conn Connection,
	dbRequest DeleteFollowsDBRequest) DeleteFollowsDBResponse {

	var dbResponse DeleteFollowsDBResponse
	dbResponse.ProjectID = dbRequest.ProjectID

	rows, err := conn.Gateway.Query(
		`DELETE FROM rewild.follows CASCADE
                                  WHERE project_id=$1;`,
		dbRequest.ProjectID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return dbResponse

}
