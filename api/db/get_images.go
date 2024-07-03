package db

type GetImagesDBRequest struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
}

type GetImagesDBResponse struct {
	Images []GetImageDBResponse `json:"images"`
}

func GetImages(
	conn Connection,
	dbRequest GetImagesDBRequest) GetImagesDBResponse {

	var dbResponse GetImagesDBResponse
	var currImageDBResponse GetImageDBResponse
	currImageDBResponse.TimelinePostID = dbRequest.TimelinePostID

	rows, err := conn.Gateway.Query(
		`SELECT image_id, alt_text FROM rewild.images WHERE timeline_post_id=$1;`,
		nullIDString(dbRequest.TimelinePostID))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&currImageDBResponse.TimelinePostID,
			&currImageDBResponse.AltText,
		)

		dbResponse.Images = append(dbResponse.Images, currImageDBResponse)

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
