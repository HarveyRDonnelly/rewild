package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func DebugHandler(c *gin.Context) {
	println(fmt.Sprintf("URL: %+v\n", c.Request.URL))
	println(fmt.Sprintf("METHOD: %s\n", c.Request.Method))
	println(fmt.Sprintf("MULTIPART: %+v\n", c.Request.MultipartForm))
	println(fmt.Sprintf("CONTENT LENGTH: %+v\n", c.Request.ContentLength))
	println(fmt.Sprintf("REQUEST URI: %+v\n", c.Request.RequestURI))

	c.Next()
}
