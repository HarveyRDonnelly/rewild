package db

type GetAuthDBRequest struct {
	UserID uuid_t `json:"user_id"`
}

type GetAuthDBResponse struct {
	UserID    uuid_t   `json:"user_id"`
	Password	string `json:"password"`
}

func GetAuth(conn Connection, dbRequest GetAuthDBRequest) GetAuthDBResponse {

	var dbResponse GetAuthDBResponse
	dbResponse.UserID = dbRequest.UserID

	rows, err := conn.Gateway.Query(
		`SELECT user_id, password FROM rewild.auths WHERE user_id=$1;`,
		dbRequest.UserID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&dbResponse.UserID,
			&dbResponse.Password,)

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
