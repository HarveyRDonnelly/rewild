package db

import (
	"github.com/google/uuid"
	"github.com/gookit/config/v2"
	"os"
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

		currTimelinePostImagesDBResponse := GetTimelinePostImages(
			conn,
			GetTimelinePostImagesDBRequest{
				TimelinePostID: currTimelinePostID,
			},
		)

		var currTimelinePostImages = make([]entities.Image, 0)

		for i := 0; i < len(currTimelinePostImagesDBResponse.Images); i++ {
			currTimelinePostImages = append(
				currTimelinePostImages,
				ConstructImage(
					conn,
					currTimelinePostImagesDBResponse.Images[i],
				),
			)
		}

		currTimelinePost := entities.TimelinePost{
			TimelinePostID: currTimelinePostDBResponse.TimelinePostID,
			NextID:         currTimelinePostDBResponse.NextID,
			PrevID:         uuid.Nil,
			Title:          currTimelinePostDBResponse.Title,
			Body:           currTimelinePostDBResponse.Body,
			Images:         currTimelinePostImages,
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

			nextTimelinePostImagesDBResponse := GetTimelinePostImages(
				conn,
				GetTimelinePostImagesDBRequest{
					TimelinePostID: nextTimelinePostID,
				},
			)

			var nextTimelinePostImages []entities.Image

			for i := 0; i < len(nextTimelinePostImagesDBResponse.Images); i++ {
				nextTimelinePostImages = append(
					nextTimelinePostImages,
					ConstructImage(
						conn,
						nextTimelinePostImagesDBResponse.Images[i],
					),
				)
			}

			nextTimelinePost := entities.TimelinePost{
				TimelinePostID: nextTimelinePostDBResponse.TimelinePostID,
				NextID:         nextTimelinePostDBResponse.TimelinePostID,
				PrevID:         currTimelinePostID,
				Title:          nextTimelinePostDBResponse.Title,
				Body:           nextTimelinePostDBResponse.Body,
				Images:         nextTimelinePostImages,
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

func ConstructDiscussionBoardMessageLimited(
	conn Connection,
	dbResponse GetDiscussionBoardMessageDBResponse,
	rootLimit int,
	depthLimit int) entities.DiscussionBoardMessage {

	var childMessages = make([]entities.DiscussionBoardMessage, 0)
	var currChildMessage entities.DiscussionBoardMessage

	currMessage := entities.DiscussionBoardMessage{
		DiscussionBoardMessageID: dbResponse.DiscussionBoardMessageID,
		Body:                     dbResponse.Body,
		AuthorID:                 dbResponse.AuthorID,
	}

	// Retrieve message children
	childMessagesDBResponse := GetDiscussionBoardMessageChildren(
		conn,
		GetDiscussionBoardMessageChildrenDBRequest{
			ParentMessageID: dbResponse.DiscussionBoardMessageID,
		})

	// rootLimit is -1 if visiting non-root node
	if rootLimit < 1 {
		rootLimit = len(childMessagesDBResponse.ChildMessages)
	}

	if len(childMessagesDBResponse.ChildMessages) < rootLimit {
		rootLimit = len(childMessagesDBResponse.ChildMessages)
	}

	// Recursively construct children
	if depthLimit > 0 {

		for i := 0; i < rootLimit; i++ {

			currChildMessage = ConstructDiscussionBoardMessageLimited(
				conn,
				childMessagesDBResponse.ChildMessages[i],
				rootLimit-1,
				depthLimit-1,
			)

			childMessages = append(
				childMessages,
				currChildMessage,
			)
		}

	}

	currMessage.Children = childMessages

	return currMessage

}

func ConstructDiscussionBoardLimited(
	conn Connection,
	dbResponse GetDiscussionBoardDBResponse,
	rootLimit int,
	depthLimit int) entities.DiscussionBoard {

	rootMessageDBResponse := GetDiscussionBoardMessage(
		conn,
		GetDiscussionBoardMessageDBRequest{
			DiscussionBoardMessageID: dbResponse.RootID,
		})

	rootMessage := ConstructDiscussionBoardMessageLimited(
		conn,
		rootMessageDBResponse,
		rootLimit,
		depthLimit)

	return entities.DiscussionBoard{
		DiscussionBoardID: dbResponse.DiscussionBoardID,
		Root:              &rootMessage,
	}

}

func ConstructProject(
	conn Connection,
	dbResponse GetProjectDBResponse) entities.Project {

	// Load project absolute path
	var absolutePath, _ = os.LookupEnv("PROJECT_PATH")

	// Load environment variables
	var whichEnv, isEnvSet = os.LookupEnv("SERVER_ENV")
	if !isEnvSet {
		whichEnv = "default"
	}

	configFileBytes, _ := os.ReadFile(absolutePath + "config/" + whichEnv + ".json")
	configFileStr := string(configFileBytes)
	configFileStr = os.ExpandEnv(configFileStr)

	// Load environment config
	err := config.LoadStrings("json", configFileStr)
	if err != nil {
		panic(err)
	}

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

	// Retrieve discussion board info
	discussionBoardDBResponse := GetDiscussionBoard(
		conn,
		GetDiscussionBoardDBRequest{
			DiscussionBoardID: dbResponse.DiscussionBoardID,
		},
	)
	discussionBoard := ConstructDiscussionBoardLimited(
		conn,
		discussionBoardDBResponse,
		config.Int("discussion_board.root_limit"),
		config.Int("discussion_board.depth_limit"),
	)

	return entities.Project{
		ProjectID:       dbResponse.ProjectID,
		Name:            dbResponse.Name,
		Description:     dbResponse.Description,
		Pindrop:         &pindrop,
		Timeline:        &timeline,
		DiscussionBoard: &discussionBoard,
		FollowerCount:   dbResponse.FollowerCount,
	}

}
