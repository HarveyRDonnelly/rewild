package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {

	print("< AUTH MIDDLEWARE PLACEHOLDER >\n")

	c.Next()
}
