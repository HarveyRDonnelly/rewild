package db

type UpdateImageDBRequest struct {
	ImageID        uuid_t `json:"image_id"`
	TimelinePostID uuid_t `json:"timeline_post_id"`
	AltText        string `json:"alt_text"`
}

type UpdateImageDBResponse struct {
	ImageID        uuid_t `json:"image_id"`
	TimelinePostID uuid_t `json:"timeline_post_id"`
	AltText        string `json:"alt_text"`
}

func UpdateImage(
	conn Connection,
	dbRequest UpdateImageDBRequest) UpdateImageDBResponse {

	rows, err := conn.Gateway.Query(
		`UPDATE rewild.images SET timeline_post_id=$2, alt_text=$3 WHERE image_id=$1;;`,
		nullIDString(dbRequest.ImageID),
		nullIDString(dbRequest.TimelinePostID),
		dbRequest.AltText,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	dbResponse := UpdateImageDBResponse(
		GetImage(
			conn,
			GetImageDBRequest{
				ImageID: dbRequest.ImageID,
			},
		),
	)

	return dbResponse

}
