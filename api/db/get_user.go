package db

type GetUserDBRequest struct {
	UserID uuid_t `json:"user_id"`
}

type GetUserDBResponse struct {
	UserID    uuid_t   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Follows   []uuid_t `json:"follows"`
}

func GetUser(conn Connection, dbRequest GetUserDBRequest) GetUserDBResponse {

	var dbResponse GetUserDBResponse
	dbResponse.UserID = dbRequest.UserID

	rows, err := conn.Gateway.Query(
		`SELECT first_name, last_name, email, username FROM rewild.users WHERE user_id=$1;`,
		dbRequest.UserID.String())

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&dbResponse.FirstName,
			&dbResponse.LastName,
			&dbResponse.Email,
			&dbResponse.Username)

		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	dbResponse.Follows = getUsersProjectIDs(conn, dbRequest.UserID)

	return dbResponse
}

func getUsersProjectIDs(conn Connection, userID uuid_t) []uuid_t {

	var currProjectID uuid_t

	projectIDs := []uuid_t{}

	rows, err := conn.Gateway.Query(
		`SELECT project_id FROM rewild.follows WHERE user_id=$1;`,
		nullIDString(userID))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&currProjectID)
		if err != nil {
			panic(err)
		}
		projectIDs = append(projectIDs, currProjectID)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return projectIDs
}
