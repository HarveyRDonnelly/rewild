package db

type CreateUserDBResponse struct {
	UserID    uuid_t   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Follows   []uuid_t `json:"follows"`
}

type CreateUserDBRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func CreateUser(
	conn Connection,
	dbRequest CreateUserDBRequest) CreateUserDBResponse {

	var userID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.users (first_name, last_name, username, email)
					VALUES ($1, $2, $3, $4) RETURNING user_id;`,
		dbRequest.FirstName,
		dbRequest.LastName,
		dbRequest.Username,
		dbRequest.Email,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	dbResponse := CreateUserDBResponse(
		GetUser(
			conn,
			GetUserDBRequest{
				UserID: userID,
			},
		),
	)

	return dbResponse

}
