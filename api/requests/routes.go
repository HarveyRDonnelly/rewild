package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"rewild-it/api/db"
)

// Alias UUID type
type uuid_t = uuid.UUID

var DB db.Connection

func SetDB(db db.Connection) {
	DB = db
}

func Create() *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20

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

	return r
}
