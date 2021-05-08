package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.DisableConsoleColor()
	gin.ForceConsoleColor()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Run()
}
