package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"rewild-it/api/db"
	"rewild-it/api/middleware"
)

// Alias UUID type
type uuid_t = uuid.NullUUID

var DB db.Connection

func SetDB(db db.Connection) {
	DB = db
}

func Create() *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 30

	// Middleware (temporarily disabled)
	r.Use(middleware.AuthHandler)
	//r.Use(middleware.DebugHandler)

	// Routes
	r = getUserRoute(r)
	r = createUserRoute(r)
	r = getProjectRoute(r)
	r = createProjectRoute(r)
	r = createTimelinePostRoute(r)
	r = updateProjectNameRoute(r)
	r = updateProjectDescriptionRoute(r)
	r = getPindropsRoute(r)
	r = createImageRoute(r)
	r = updateImageAltTextRoute(r)
	r = createTimelinePostImageRoute(r)
	r = createDiscussionBoardMessageRoute(r)
	r = createFollowRoute(r)
	r = deleteFollowRoute(r)
	r = deleteDiscussionBoardMessageRoute(r)
	r = deleteTimelinePostRoute(r)
	r = deleteProjectRoute(r)
	r = createUserSessionRoute(r)
	r = getUserSessionRoute(r)

	// Load project absolute path
	var absolutePath, _ = os.LookupEnv("PROJECT_PATH")

	r.Static("/images/files/", absolutePath+"/res")

	return r
}
