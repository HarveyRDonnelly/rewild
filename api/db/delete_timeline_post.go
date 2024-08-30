package db

type DeleteTimelinePostDBRequest struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
}

type DeleteTimelinePostDBResponse struct {
	TimelinePostID uuid_t `json:"timeline_post_id"`
}

func DeleteTimelinePost(
	conn Connection,
	dbRequest DeleteTimelinePostDBRequest) DeleteTimelinePostDBResponse {

	var dbResponse DeleteTimelinePostDBResponse
	dbResponse.TimelinePostID = dbRequest.TimelinePostID

	rows, err := conn.Gateway.Query(
		`DELETE FROM rewild.timeline_posts
                WHERE timeline_post_id=$1;`,
		nullIDString(dbRequest.TimelinePostID))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return dbResponse

}
