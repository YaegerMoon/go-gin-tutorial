package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<h1>hello, world</h1>",
		})
	})

	r.GET("/pureJSON", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<h1>hello, world</h1>",
		})
	})

	r.Run(":8080")
}
