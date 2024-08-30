package db

type GetTimelinePostImagesDBRequest struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
}

type GetTimelinePostImagesDBResponse struct {
	Images []GetImageDBResponse `json:"images"`
}

func GetTimelinePostImages(
	conn Connection,
	dbRequest GetTimelinePostImagesDBRequest) GetTimelinePostImagesDBResponse {

	var dbResponse GetTimelinePostImagesDBResponse
	var currImageDBResponse GetImageDBResponse
	var currImageDBRequest GetImageDBRequest

	rows, err := conn.Gateway.Query(
		`SELECT image_id FROM rewild.timeline_post_images WHERE timeline_post_id=$1
				ORDER BY arr_index;`,
		dbRequest.TimelinePostID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&currImageDBRequest.ImageID,
		)

		currImageDBResponse = GetImage(
			conn,
			currImageDBRequest,
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
