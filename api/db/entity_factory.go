package db

import (
	"github.com/google/uuid"
	"rewild-it/api/entities"
)

func ConstructPindrop(
	_ Connection,
	dbResponse GetPindropDBResponse) entities.Pindrop {

	return entities.Pindrop{
		PindropID: dbResponse.PindropID,
		Longitude: dbResponse.Longitude,
		Latitude:  dbResponse.Latitude,
	}
}

func ConstructTimeline(
	conn Connection,
	dbResponse GetTimelineDBResponse) entities.Timeline {

	var timeline entities.Timeline
	timeline.TimelineID = dbResponse.TimelineID

	if dbResponse.HeadID != uuid.Nil {

		currTimelinePostID := dbResponse.HeadID

		currTimelinePostDBResponse := GetTimelinePost(
			conn,
			GetTimelinePostDBRequest{
				TimelinePostID: currTimelinePostID,
			},
		)

		nextTimelinePostID := currTimelinePostDBResponse.NextID

		currTimelinePost := &entities.TimelinePost{
			TimelinePostID: currTimelinePostDBResponse.TimelinePostID,
			Next:           nil,
			Prev:           nil,
			Title:          currTimelinePostDBResponse.Title,
			Body:           currTimelinePostDBResponse.Body,
		}

		timeline.Head = currTimelinePost

		for nextTimelinePostID != uuid.Nil {

			nextTimelinePostDBResponse := GetTimelinePost(
				conn,
				GetTimelinePostDBRequest{
					TimelinePostID: nextTimelinePostID,
				},
			)

			nextTimelinePost := &entities.TimelinePost{
				TimelinePostID: nextTimelinePostDBResponse.TimelinePostID,
				Next:           nil,
				Prev:           currTimelinePost,
				Title:          nextTimelinePostDBResponse.Title,
				Body:           nextTimelinePostDBResponse.Body,
			}

			currTimelinePostID = nextTimelinePostDBResponse.TimelinePostID
			nextTimelinePostID = nextTimelinePostDBResponse.NextID
			currTimelinePost = nextTimelinePost

		}

		timeline.Tail = currTimelinePost

	} else {
		timeline.Head = nil
		timeline.Tail = nil
	}
	return timeline

}

func ConstructProject(
	conn Connection,
	dbResponse GetProjectDBResponse) entities.Project {

	// Retrieve pindrop info
	pindropDBResponse := GetPindrop(
		conn,
		GetPindropDBRequest{
			PindropID: dbResponse.PindropID,
		},
	)
	pindrop := ConstructPindrop(conn, pindropDBResponse)

	// Retrieve timeline info
	timelineDBResponse := GetTimeline(
		conn,
		GetTimelineDBRequest{
			TimelineID: dbResponse.TimelineID,
		},
	)
	timeline := ConstructTimeline(conn, timelineDBResponse)

	return entities.Project{
		ProjectID:     dbResponse.ProjectID,
		Name:          dbResponse.Name,
		Description:   dbResponse.Description,
		Pindrop:       &pindrop,
		Timeline:      &timeline,
		FollowerCount: dbResponse.FollowerCount,
	}

}
