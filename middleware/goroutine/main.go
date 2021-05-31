package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in routine", cCp.Request.URL.Path)

		}()

		log.Println("Done! in path", c.Request.URL.Path)
		c.Status(http.StatusOK)
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done in path", c.Request.URL.Path)
		c.Status(http.StatusOK)
	})

	r.Run(":8080")
}
