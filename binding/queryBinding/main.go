package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()
	r.Any("/testing", func(c *gin.Context) {
		var person Person
		if c.ShouldBindQuery(&person) == nil {
			log.Println(person.Name)
			log.Println(person.Address)
		}
		c.String(http.StatusOK, "Success")
	})
	r.Run(":8080")
}
