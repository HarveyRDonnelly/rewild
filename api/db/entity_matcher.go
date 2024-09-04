package db

func FindProjectIDByPindropID(conn Connection, pindropID uuid_t) uuid_t {

	var projectID uuid_t

	rows, err := conn.Gateway.Query(
		`SELECT project_id FROM rewild.projects WHERE pindrop_id=$1`,
		pindropID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&projectID)

		if err != nil {
			panic(err)
		}
	}

	return projectID

}

func FindUserIDByUsername(conn Connection, username string) uuid_t {

	var userID uuid_t

	rows, err := conn.Gateway.Query(
		`SELECT user_id FROM rewild.users WHERE username=$1`,
		userID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&userID)

		if err != nil {
			panic(err)
		}
	}

	return userID

}