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

		currTimelinePost := entities.TimelinePost{
			TimelinePostID: currTimelinePostDBResponse.TimelinePostID,
			NextID:         currTimelinePostDBResponse.NextID,
			PrevID:         uuid.Nil,
			Title:          currTimelinePostDBResponse.Title,
			Body:           currTimelinePostDBResponse.Body,
		}

		timeline.Posts = append(timeline.Posts, currTimelinePost)

		timeline.HeadID = currTimelinePostID

		for nextTimelinePostID != uuid.Nil {

			nextTimelinePostDBResponse := GetTimelinePost(
				conn,
				GetTimelinePostDBRequest{
					TimelinePostID: nextTimelinePostID,
				},
			)

			nextTimelinePost := entities.TimelinePost{
				TimelinePostID: nextTimelinePostDBResponse.TimelinePostID,
				NextID:         nextTimelinePostDBResponse.TimelinePostID,
				PrevID:         currTimelinePostID,
				Title:          nextTimelinePostDBResponse.Title,
				Body:           nextTimelinePostDBResponse.Body,
			}

			timeline.Posts = append(timeline.Posts, nextTimelinePost)

			currTimelinePostID = nextTimelinePostDBResponse.TimelinePostID
			nextTimelinePostID = nextTimelinePostDBResponse.NextID
			currTimelinePost = nextTimelinePost

		}

		timeline.TailID = currTimelinePostID

	} else {
		timeline.HeadID = uuid.Nil
		timeline.TailID = uuid.Nil
	}
	return timeline

}

func ConstructImage(
	_ Connection,
	dbResponse GetImageDBResponse) entities.Image {
	return entities.Image{
		ImageID: dbResponse.ImageID,
		AltText: dbResponse.AltText,
	}
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
