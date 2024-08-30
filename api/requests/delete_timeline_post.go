package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rewild-it/api/db"
)

func deleteTimelinePostRoute(r *gin.Engine) *gin.Engine {

	r.DELETE("/project/:project_id/timeline/post/:timeline_post_id", func(c *gin.Context) {

		var timelinePostID = uuid.Must(uuid.Parse(c.Param("timeline_post_id")))
		var projectID = uuid.Must(uuid.Parse(c.Param("project_id")))

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

		if currTimelinePostDBResponse.PrevID == uuid.Nil {
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
				db.UpdateTimelinePostDBRequest(prevTimelinePostDBResponse),
			)
		}

		if currTimelinePostDBResponse.NextID == uuid.Nil {
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
				db.UpdateTimelinePostDBRequest(nextTimelinePostDBResponse),
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
