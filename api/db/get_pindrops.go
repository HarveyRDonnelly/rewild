package db

type GetPindropsDBRequest struct {
	Delta           float64 `json:"delta"`
	CentreLongitude float64 `json:"centre_longitude"`
	CentreLatitude  float64 `json:"centre_latitude"`
}

type GetPindropsDBResponse struct {
	Pindrops []GetPindropDBResponse
}

func GetPindrops(
	conn Connection,
	dbRequest GetPindropsDBRequest) GetPindropsDBResponse {

	var dbResponse GetPindropsDBResponse
	var currPindropDBResponse GetPindropDBResponse

	rows, err := conn.Gateway.Query(
		`SELECT pindrop_id, latitude, longitude FROM rewild.pindrops 
                    WHERE ($2::FLOAT - $1::FLOAT) <= longitude AND longitude <= ($2::FLOAT + $1::FLOAT)
					AND   ($3::FLOAT - $1::FLOAT) <= latitude AND latitude <= ($3::FLOAT + $1::FLOAT);`,
		dbRequest.Delta,
		dbRequest.CentreLongitude,
		dbRequest.CentreLatitude,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&currPindropDBResponse.PindropID,
			&currPindropDBResponse.Latitude,
			&currPindropDBResponse.Longitude,
		)

		dbResponse.Pindrops = append(dbResponse.Pindrops, currPindropDBResponse)

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
