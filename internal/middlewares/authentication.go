package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate authentication here
		// if authentication fails:
		// c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		// return

		// if authentication succeeds:
		fmt.Printf("mock jwt to test")
		c.Next()
	}
}
