package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TrailingSlashMiddleware redirect to the URL without the trailing slash.
// Convention: URLs without trailing slash.
func TrailingSlashMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path[len(c.Request.URL.Path)-1] == '/' {
			c.Redirect(http.StatusMovedPermanently, c.Request.URL.Path[:len(c.Request.URL.Path)-1])
			c.Abort()
			return
		}
		c.Next()
	}
}
