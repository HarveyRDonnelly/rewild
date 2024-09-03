package db

type DeletePindropDBRequest struct {
	PindropID uuid_t `json:"pindrop_id"`
}

type DeletePindropDBResponse struct {
	PindropID uuid_t `json:"pindrop_id"`
}

func DeletePindrop(
	conn Connection,
	dbRequest DeletePindropDBRequest) DeletePindropDBResponse {

	var dbResponse DeletePindropDBResponse
	dbResponse.PindropID = dbRequest.PindropID

	rows, err := conn.Gateway.Query(
		`DELETE FROM rewild.pindrops CASCADE
                                  WHERE pindrop_id=$1;`,
		dbRequest.PindropID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return dbResponse

}
