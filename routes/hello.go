package routes

import (
	"gin-api-boilerplate/api"
	"github.com/gin-gonic/gin"
)

// Hello ...
func Hello(r *gin.Engine) {
	r.GET("/hello", api.Hello)
}
