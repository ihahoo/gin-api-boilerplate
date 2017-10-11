package main

import (
	"fmt"

	"gin-api-boilerplate/routes"

	"github.com/gin-gonic/gin"
	"github.com/ihahoo/go-api-lib/config"
	"github.com/ihahoo/go-api-lib/log"
)

// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	name := config.GetString("name")
	port := config.GetString("port")
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.Use(CORSMiddleware())

	routes.Routes(r)

	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"errcode": 405, "errmsg": "Method Not Allowed"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"errcode": 404, "errmsg": "Not Found"})
	})

	log.GetLog().Info(name + " start listening at http://loaclhost:" + port)
	fmt.Printf("==> 🚀 %s listening at %s\n", name, port)
	r.Run(":" + port)
}
