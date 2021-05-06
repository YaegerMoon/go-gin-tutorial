package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"method":  "GET",
			"urlpath": "/",
		})
	})

	v1.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
			"url":    c.Request.URL,
		})
	})
	v1.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
			"url":    c.Request.URL,
		})
	})
	v1.GET("/c", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
			"url":    c.Request.URL,
		})
	})

	v2 := router.Group("/v2")
	{
		v2.GET("/a", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
				"url":    c.Request.URL,
			})
		})
		v2.GET("/b", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
				"url":    c.Request.URL,
			})
		})
		v2.GET("/c", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
				"url":    c.Request.URL,
			})
		})
	}

	router.Run(":8080")

}
