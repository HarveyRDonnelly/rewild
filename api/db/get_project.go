package db

type GetProjectResponse struct {
	ProjectID         uuid_t `json:"project_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	PindropID         uuid_t `json:"pindrop_id"`
	TimelineID        uuid_t `json:"timeline_id"`
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	FollowerCount     int    `json:"follower_count"`
}

func GetProject(conn Connection, projectID uuid_t) GetProjectResponse {

	var name string
	var description string
	var pindropID uuid_t
	var timelineID uuid_t
	var discussionBoardID uuid_t
	var followerCount int

	// Get user information
	rows, err := conn.Gateway.Query(
		`SELECT name, description, pindrop_id, timeline_id, discussion_board_id, follower_count 
				FROM rewild.projects WHERE project_id=$1;`,
		projectID.String())

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &description, &pindropID, &timelineID, &discussionBoardID, &followerCount)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	project := GetProjectResponse{
		ProjectID:         projectID,
		Name:              name,
		Description:       description,
		PindropID:         pindropID,
		TimelineID:        timelineID,
		DiscussionBoardID: discussionBoardID,
		FollowerCount:     followerCount,
	}

	return project
}
