package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

func deleteTimelinePostRoute(r *gin.Engine) *gin.Engine {

	r.DELETE("/project/:project_id/timeline/post/:timeline_post_id", func(c *gin.Context) {

		var timelinePostID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("timeline_post_id"))),
			Valid: true,
		}
		var projectID = uuid.NullUUID{
			UUID:  uuid.Must(uuid.Parse(c.Param("project_id"))),
			Valid: true,
		}

		projectDBResponse := db.GetProject(
			DB,
			db.GetProjectDBRequest{
				ProjectID: projectID,
			},
		)

		currProject := db.ConstructProject(DB, projectDBResponse)
		currTimeline := currProject.Timeline

		currTimelineDBResponse := db.GetTimeline(
			DB,
			db.GetTimelineDBRequest{
				TimelineID: currTimeline.TimelineID,
			},
		)

		currTimelinePostDBResponse := db.GetTimelinePost(
			DB,
			db.GetTimelinePostDBRequest{
				TimelinePostID: timelinePostID,
			},
		)

		if currTimelinePostDBResponse.PrevID.Valid == false {
			currTimelineDBResponse.HeadID = currTimelinePostDBResponse.NextID
		} else {
			prevTimelinePostDBResponse := db.GetTimelinePost(
				DB,
				db.GetTimelinePostDBRequest{
					TimelinePostID: currTimelinePostDBResponse.PrevID,
				},
			)
			prevTimelinePostDBResponse.NextID = currTimelinePostDBResponse.NextID
			db.UpdateTimelinePost(
				DB,
				db.UpdateTimelinePostDBRequest{
					TimelinePostID: prevTimelinePostDBResponse.TimelinePostID,
					NextID:         prevTimelinePostDBResponse.NextID,
					PrevID:         prevTimelinePostDBResponse.PrevID,
					Title:          prevTimelinePostDBResponse.Title,
					Body:           prevTimelinePostDBResponse.Body,
					Type:           prevTimelinePostDBResponse.Type,
					AuthorID:       prevTimelinePostDBResponse.AuthorID,
				},
			)
		}

		if currTimelinePostDBResponse.NextID.Valid == false {
			currTimelineDBResponse.TailID = currTimelinePostDBResponse.PrevID
		} else {
			nextTimelinePostDBResponse := db.GetTimelinePost(
				DB,
				db.GetTimelinePostDBRequest{
					TimelinePostID: currTimelinePostDBResponse.NextID,
				},
			)
			nextTimelinePostDBResponse.PrevID = currTimelinePostDBResponse.PrevID
			db.UpdateTimelinePost(
				DB,
				db.UpdateTimelinePostDBRequest{
					TimelinePostID: nextTimelinePostDBResponse.TimelinePostID,
					NextID:         nextTimelinePostDBResponse.NextID,
					PrevID:         nextTimelinePostDBResponse.PrevID,
					Title:          nextTimelinePostDBResponse.Title,
					Body:           nextTimelinePostDBResponse.Body,
					Type:           nextTimelinePostDBResponse.Type,
					AuthorID:       nextTimelinePostDBResponse.AuthorID,
				},
			)
		}

		db.UpdateTimeline(
			DB,
			db.UpdateTimelineDBRequest(currTimelineDBResponse),
		)

		db.DeleteTimelinePost(
			DB,
			db.DeleteTimelinePostDBRequest{
				TimelinePostID: timelinePostID,
			},
		)

		c.Status(http.StatusOK)
	})

	return r

}
