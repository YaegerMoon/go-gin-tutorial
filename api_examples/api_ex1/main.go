package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/someget", getting)
	router.POST("/somepost", posting)
	router.OPTIONS("/someoptions", options)
	router.DELETE("/somedelete", deleting)
	router.PATCH("/somepatch", patching)
	router.HEAD("/somehead", head)

	router.Run(":3000")
}

func getting(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "get",
	})
}

func posting(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "post",
	})
}

func options(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "option",
	})
}

func deleting(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "delete",
	})
}

func patching(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "patch",
	})
}

func head(c *gin.Context) {
	fmt.Print("HEAD")
	c.JSON(200, gin.H{
		"method": "head",
	})
}
