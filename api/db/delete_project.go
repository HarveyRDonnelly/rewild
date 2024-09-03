package db

type DeleteProjectDBRequest struct {
	ProjectID uuid_t `json:"project_id"`
}

type DeleteProjectDBResponse struct {
	ProjectID uuid_t `json:"project_id"`
}

func DeleteProject(
	conn Connection,
	dbRequest DeleteProjectDBRequest) DeleteProjectDBResponse {

	var dbResponse DeleteProjectDBResponse
	dbResponse.ProjectID = dbRequest.ProjectID

	rows, err := conn.Gateway.Query(
		`DELETE FROM rewild.projects CASCADE
                                  WHERE project_id=$1;`,
		dbRequest.ProjectID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return dbResponse

}
