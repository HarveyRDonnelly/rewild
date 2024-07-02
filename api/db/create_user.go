package db

type CreateUserResponse struct {
	UserID    uuid_t   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Follows   []uuid_t `json:"follows"`
}

func CreateUser(
	conn Connection,
	firstName string,
	lastName string,
	username string,
	email string) CreateUserResponse {

	var userID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.users (first_name, last_name, username, email)
					VALUES ($1, $2, $3, $4) RETURNING user_id;`,
		firstName,
		lastName,
		username,
		email,
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

	user := CreateUserResponse(GetUser(conn, userID))

	return user

}
