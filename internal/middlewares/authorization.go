package middlewares

import "github.com/gin-gonic/gin"

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate authorization here
		// if authorization fails:
		// c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
		// return

		// if authorization succeeds:
		c.Next()
	}
}
