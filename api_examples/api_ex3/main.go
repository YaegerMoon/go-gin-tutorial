package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TEST : localhost:8080/welcome?lastname=Mun

func main() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

	})
	router.Run(":8080")
}
