package db

type UpdateProjectDBRequest struct {
	ProjectID         uuid_t `json:"project_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	PindropID         uuid_t `json:"pindrop_id"`
	TimelineID        uuid_t `json:"timeline_id"`
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	FollowerCount     int    `json:"follower_count"`
}

type UpdateProjectDBResponse struct {
	ProjectID         uuid_t `json:"project_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	PindropID         uuid_t `json:"pindrop_id"`
	TimelineID        uuid_t `json:"timeline_id"`
	DiscussionBoardID uuid_t `json:"discussion_board_id"`
	FollowerCount     int    `json:"follower_count"`
}

func UpdateProject(
	conn Connection,
	dbRequest UpdateProjectDBRequest) UpdateProjectDBResponse {

	rows, err := conn.Gateway.Query(
		`UPDATE rewild.projects 
			SET name=$2, description=$3, pindrop_id=$4, timeline_id=$5, discussion_board_id=$6, follower_count=$7 
				WHERE project_id=$1;;`,
		nullIDString(dbRequest.ProjectID),
		dbRequest.Name,
		dbRequest.Description,
		nullIDString(dbRequest.PindropID),
		nullIDString(dbRequest.TimelineID),
		nullIDString(dbRequest.DiscussionBoardID),
		dbRequest.FollowerCount,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	dbResponse := UpdateProjectDBResponse(
		GetProject(
			conn,
			GetProjectDBRequest{
				ProjectID: dbRequest.ProjectID,
			},
		),
	)

	return dbResponse

}
