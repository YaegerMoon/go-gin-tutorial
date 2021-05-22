package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.StaticFS("/more_static", http.Dir("./assets"))
	r.StaticFile("/a", "./assets/a.txt")
	r.StaticFile("/a.png", "./assets/a.png")

	r.Run(":8080")
}
