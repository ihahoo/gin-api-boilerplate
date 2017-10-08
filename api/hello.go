package api

import (
	"github.com/gin-gonic/gin"
)

// Hello ...
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "hello world!",
	})
}
