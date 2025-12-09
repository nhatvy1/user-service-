package middlewares

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	// Implement CORS middleware logic here
	return func(c *gin.Context) {
		c.Next()
	}
}
