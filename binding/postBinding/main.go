package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", startPage)
	r.POST("/testing", startPage)
	r.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person

	//"GET Method 이면 Query Binding"
	//"POST Method 이면 JSON, XML Binding"
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}
	c.String(http.StatusOK, "Succces")
}
