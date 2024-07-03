package db

type UpdateTimelineDBRequest struct {
	TimelineID uuid_t `json:"timeline_post_id"`
	HeadID     uuid_t `json:"head_id"`
	TailID     uuid_t `json:"tail_id"`
}

type UpdateTimelineDBResponse struct {
	TimelineID uuid_t `json:"timeline_post_id"`
	HeadID     uuid_t `json:"head_id"`
	TailID     uuid_t `json:"tail_id"`
}

func UpdateTimeline(
	conn Connection,
	dbRequest UpdateTimelineDBRequest) UpdateTimelineDBResponse {

	rows, err := conn.Gateway.Query(
		`UPDATE rewild.timelines SET head_id=$2, tail_id=$3 WHERE timeline_id=$1;;`,
		nullIDString(dbRequest.TimelineID),
		nullIDString(dbRequest.HeadID),
		nullIDString(dbRequest.TailID),
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	dbResponse := UpdateTimelineDBResponse(
		GetTimeline(
			conn,
			GetTimelineDBRequest{
				TimelineID: dbRequest.TimelineID,
			},
		),
	)

	return dbResponse

}
