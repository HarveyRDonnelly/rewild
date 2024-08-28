package db

type CreateImageDBRequest struct {
	AltText string `json:"alt_text"`
}

type CreateImageDBResponse struct {
	ImageID uuid_t `json:"image_id"`
	AltText string `json:"alt_text"`
}

func CreateImage(
	conn Connection,
	dbRequest CreateImageDBRequest) CreateImageDBResponse {

	var imageID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.images (alt_text)
				VALUES ($1) RETURNING image_id;`,
		dbRequest.AltText,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&imageID)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	dbResponse := CreateImageDBResponse(
		GetImage(
			conn,
			GetImageDBRequest{
				ImageID: imageID,
			},
		),
	)

	return dbResponse

}
