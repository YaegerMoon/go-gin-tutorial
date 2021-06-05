package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.TrustedProxies = []string{"127.0.0.1"}

	router.GET("/", func(c *gin.Context) {

		fmt.Printf("ClientIP: %s\n", c.ClientIP())
	})
	router.Run(":8080")

}
