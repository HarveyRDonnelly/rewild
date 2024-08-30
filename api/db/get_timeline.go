package db

type GetTimelineDBRequest struct {
	TimelineID uuid_t `json:"timeline_id"`
}

type GetTimelineDBResponse struct {
	TimelineID uuid_t `json:"timeline_id"`
	HeadID     uuid_t `json:"head_id"`
	TailID     uuid_t `json:"tail_id"`
}

func GetTimeline(
	conn Connection,
	dbRequest GetTimelineDBRequest) GetTimelineDBResponse {

	var dbResponse GetTimelineDBResponse
	dbResponse.TimelineID = dbRequest.TimelineID

	rows, err := conn.Gateway.Query(
		`SELECT head_id, tail_id FROM rewild.timelines WHERE timeline_id=$1;`,
		dbRequest.TimelineID,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&dbResponse.HeadID,
			&dbResponse.TailID,
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
