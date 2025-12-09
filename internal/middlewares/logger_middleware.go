package middlewares

import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	// Implement logging logic here
	return func(c *gin.Context) {
		c.Next()
	}
}
