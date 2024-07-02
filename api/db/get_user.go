package db

type GetUserResponse struct {
	UserID    uuid_t   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Follows   []uuid_t `json:"follows"`
}

func GetUser(conn Connection, userID uuid_t) GetUserResponse {

	var firstName string
	var lastName string
	var username string
	var email string
	var follows []uuid_t

	// Get user information
	rows, err := conn.Gateway.Query(
		`SELECT first_name, last_name, username, email FROM rewild.users WHERE user_id=$1;`,
		userID.String())

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&firstName, &lastName, &username, &email)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// Get project follows
	follows = getUsersProjectIDs(conn, userID)

	user := GetUserResponse{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Follows:   follows,
	}

	return user
}

func getUsersProjectIDs(conn Connection, userID uuid_t) []uuid_t {

	var projectIDs []uuid_t
	var currProjectID uuid_t

	rows, err := conn.Gateway.Query(
		`SELECT project_id FROM rewild.follows WHERE user_id=$1;`,
		userID.String())

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
