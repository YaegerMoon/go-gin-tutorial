package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		if err, ok := err.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}))

	r.GET("/panic", func(c *gin.Context) {
		panic("foo")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})

	r.GET("/loop", func(c *gin.Context) {
		for {
			fmt.Print("looping")
			time.Sleep(1 * time.Second)
		}
	})

	r.Run(":8080")
}
