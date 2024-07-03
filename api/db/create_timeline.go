package db

type CreateTimelineDBRequest struct {
	HeadID uuid_t `json:"head_id"`
	TailID uuid_t `json:"tail_id"`
}

type CreateTimelineDBResponse struct {
	TimelineID uuid_t `json:"timeline_id"`
	HeadID     uuid_t `json:"head_id"`
	TailID     uuid_t `json:"tail_id"`
}

func createTimeline(
	conn Connection,
	dbRequest CreateTimelineDBRequest) CreateTimelineDBResponse {

	var timelineID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.timelines (head_id, tail_id)
				VALUES ($1, $2) RETURNING timeline_id;`,
		dbRequest.HeadID,
		dbRequest.TailID,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&timelineID)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	dbResponse := CreateTimelineDBResponse(
		GetTimeline(
			conn,
			GetTimelineDBRequest{
				TimelineID: timelineID,
			},
		),
	)

	return dbResponse

}
