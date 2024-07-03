package db

type CreateImageDBRequest struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
	AltText        string `json:"alt_text"`
}

type CreateImageDBResponse struct {
	ImageID        uuid_t `json:"image_id"`
	TimelinePostID uuid_t `json:"timeline_post_id"`
	AltText        string `json:"alt_text"`
}

func CreateImage(
	conn Connection,
	dbRequest CreateImageDBRequest) CreateImageDBResponse {

	var imageID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.images (timeline_post_id, alt_text)
				VALUES ($1) RETURNING image_id;`,
		nullIDString(dbRequest.TimelinePostID),
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
