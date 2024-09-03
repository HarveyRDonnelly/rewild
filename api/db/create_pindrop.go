package db

type CreatePindropDBRequest struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type CreatePindropDBResponse struct {
	PindropID uuid_t  `json:"pindrop_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func CreatePindrop(
	conn Connection,
	dbRequest CreatePindropDBRequest) CreatePindropDBResponse {

	var dbResponse CreatePindropDBResponse
	var pindropID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.pindrops (latitude, longitude)
		VALUES ($1, $2) RETURNING pindrop_id;`,
		dbRequest.Latitude,
		dbRequest.Longitude,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&pindropID)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	dbResponse.PindropID = pindropID
	dbResponse.Longitude = dbRequest.Longitude
	dbResponse.Latitude = dbRequest.Latitude

	return dbResponse

}
