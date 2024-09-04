package db

type CreateAuthDBRequest struct {
	UserID uuid_t `json:"user_id"`
	Password  string `json:"password"`
}

type CreateAuthDBResponse struct {
	UserID uuid_t `json:"user_id"`
}

func CreateAuth(
	conn Connection,
	dbRequest CreateAuthDBRequest) CreateAuthDBResponse {

	var dbResponse CreateAuthDBResponse
	dbResponse.UserID = dbRequest.UserID

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.auths (user_id, password)
		VALUES ($1, $2);`,
		dbRequest.UserID,
		dbRequest.Password,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return dbResponse

}
