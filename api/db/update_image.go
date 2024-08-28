package db

type UpdateImageDBRequest struct {
	ImageID uuid_t `json:"image_id"`
	AltText string `json:"alt_text"`
}

type UpdateImageDBResponse struct {
	ImageID uuid_t `json:"image_id"`
	AltText string `json:"alt_text"`
}

func UpdateImage(
	conn Connection,
	dbRequest UpdateImageDBRequest) UpdateImageDBResponse {

	rows, err := conn.Gateway.Query(
		`UPDATE rewild.images SET alt_text=$3 WHERE image_id=$1;;`,
		nullIDString(dbRequest.ImageID),
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
