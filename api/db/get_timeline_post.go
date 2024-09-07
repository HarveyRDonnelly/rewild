package db

type GetTimelinePostDBRequest struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
}

type GetTimelinePostDBResponse struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
	NextID         uuid_t `json:"next_id"`
	PrevID         uuid_t `json:"prev_id"`
	Title          string `json:"title"`
	Body           string `json:"body"`
	Type           string `json:"type"`
	AuthorID       uuid_t `json:"author_id"`
	CreatedTS      string `json:"created_ts"`
}

func GetTimelinePost(
	conn Connection,
	dbRequest GetTimelinePostDBRequest) GetTimelinePostDBResponse {

	var dbResponse GetTimelinePostDBResponse
	dbResponse.TimelinePostID = dbRequest.TimelinePostID

	rows, err := conn.Gateway.Query(
		`SELECT next_id, prev_id, title, body, type, author_id, created_ts FROM rewild.timeline_posts WHERE timeline_post_id=$1;`,
		dbRequest.TimelinePostID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&dbResponse.NextID,
			&dbResponse.PrevID,
			&dbResponse.Title,
			&dbResponse.Body,
			&dbResponse.Type,
			&dbResponse.AuthorID,
			&dbResponse.CreatedTS,
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
