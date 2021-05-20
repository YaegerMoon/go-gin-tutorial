package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		})
	})

	r.Run(":8080")
}
