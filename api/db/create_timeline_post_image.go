package db

type CreateTimelinePostImageDBRequest struct {
	ImageID        uuid_t `json:"image_id"`
	TimelinePostID uuid_t `json:"timeline_post_id"`
	ArrIndex       int    `json:"arr_index"`
}

type CreateTimelinePostImageDBResponse struct {
	ImageID uuid_t `json:"image_id"`
	AltText string `json:"alt_text"`
}

func CreateTimelinePostImage(
	conn Connection,
	dbRequest CreateTimelinePostImageDBRequest) CreateTimelinePostImageDBResponse {

	var imageID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.timeline_post_images (image_id, timeline_post_id, arr_index)
				VALUES ($1, $2, $3) RETURNING image_id;`,
		dbRequest.ImageID,
		dbRequest.TimelinePostID,
		dbRequest.ArrIndex,
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

	dbResponse := CreateTimelinePostImageDBResponse(
		GetImage(
			conn,
			GetImageDBRequest{
				ImageID: imageID,
			},
		),
	)

	return dbResponse

}
