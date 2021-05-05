package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// test : localhost:8080/form_post

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		msg := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": msg,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
