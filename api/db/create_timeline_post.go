package db

type CreateTimelinePostDBRequest struct {
	NextID uuid_t `json:"next_id"`
	PrevID uuid_t `json:"prev_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type CreateTimelinePostDBResponse struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
	NextID         uuid_t `json:"next_id"`
	PrevID         uuid_t `json:"prev_id"`
	Title          string `json:"title"`
	Body           string `json:"body"`
}

func CreateTimelinePost(
	conn Connection,
	dbRequest CreateTimelinePostDBRequest) CreateTimelinePostDBResponse {

	var timelinePostID uuid_t

	rows, err := conn.Gateway.Query(
		`INSERT INTO rewild.timeline_posts (next_id, prev_id, title, body)
				VALUES ($1, $2, $3, $4) RETURNING timeline_post_id;`,
		dbRequest.NextID,
		dbRequest.PrevID,
		dbRequest.Title,
		dbRequest.Body,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&timelinePostID)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	dbResponse := CreateTimelinePostDBResponse(
		GetTimelinePost(
			conn,
			GetTimelinePostDBRequest{
				TimelinePostID: timelinePostID,
			},
		),
	)

	return dbResponse

}
