package auth

import "github.com/gin-gonic/gin"

func Health(c *gin.Context) {
	c.Status(200)
	return
}
