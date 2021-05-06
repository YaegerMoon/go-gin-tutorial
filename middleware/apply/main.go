package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/panic", func(c *gin.Context) {
		panic("foo")
	})

	r.GET("/", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"method": "GET",
			"path":   c.Request.URL.Path,
		})
	})

	r.Run(":8080")

}
