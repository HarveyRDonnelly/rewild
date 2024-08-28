package db

type GetImageDBRequest struct {
	ImageID uuid_t `json:"image_id"`
}

type GetImageDBResponse struct {
	ImageID uuid_t `json:"image_id"`
	AltText string `json:"alt_text"`
}

func GetImage(
	conn Connection,
	dbRequest GetImageDBRequest) GetImageDBResponse {

	var dbResponse GetImageDBResponse
	dbResponse.ImageID = dbRequest.ImageID

	rows, err := conn.Gateway.Query(
		`SELECT alt_text FROM rewild.images WHERE image_id=$1;`,
		nullIDString(dbRequest.ImageID))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&dbResponse.AltText,
		)
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
