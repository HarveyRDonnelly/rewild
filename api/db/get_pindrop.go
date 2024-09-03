package db

type GetPindropDBRequest struct {
	PindropID uuid_t `json:"pindrop_id"`
}

type GetPindropDBResponse struct {
	PindropID uuid_t  `json:"pindrop_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	ProjectID uuid_t `json: project_id`
}

func GetPindrop(
	conn Connection,
	dbRequest GetPindropDBRequest) GetPindropDBResponse {

	var dbResponse GetPindropDBResponse
	dbResponse.PindropID = dbRequest.PindropID

	rows, err := conn.Gateway.Query(
		`SELECT latitude, longitude FROM rewild.pindrops WHERE pindrop_id=$1;`,
		dbRequest.PindropID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&dbResponse.Latitude,
			&dbResponse.Longitude,
		)
		if err != nil {
			panic(err)
		}
	}

	dbResponse.ProjectID = FindProjectIDByPindropID(conn, dbRequest.PindropID)

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return dbResponse

}
