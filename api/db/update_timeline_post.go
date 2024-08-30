package db

type UpdateTimelinePostDBRequest struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
	NextID         uuid_t `json:"next_id"`
	PrevID         uuid_t `json:"prev_id"`
	Title          string `json:"title"`
	Body           string `json:"body"`
}

type UpdateTimelinePostDBResponse struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
	NextID         uuid_t `json:"next_id"`
	PrevID         uuid_t `json:"prev_id"`
	Title          string `json:"title"`
	Body           string `json:"body"`
}

func UpdateTimelinePost(
	conn Connection,
	dbRequest UpdateTimelinePostDBRequest) UpdateTimelinePostDBResponse {

	rows, err := conn.Gateway.Query(
		`UPDATE rewild.timeline_posts SET next_id=$2, prev_id=$3, title=$4, body=$5 WHERE timeline_post_id=$1;;`,
		dbRequest.TimelinePostID,
		dbRequest.NextID,
		dbRequest.PrevID,
		dbRequest.Title,
		dbRequest.Body,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	dbResponse := UpdateTimelinePostDBResponse(
		GetTimelinePost(
			conn,
			GetTimelinePostDBRequest{
				TimelinePostID: dbRequest.TimelinePostID,
			},
		),
	)

	return dbResponse

}
