package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//$ curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
//$ curl -v localhost:8088/thinkerou/not-uuid

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {

	r := gin.Default()

	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})

	r.Run(":8080")
}
