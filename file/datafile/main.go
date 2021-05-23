package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/local/file", func(c *gin.Context) {
		c.File("./main.go")
	})
	f := os.DirFS("./")
	var fileSys http.FileSystem = http.FS(f)
	r.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("./main.go", fileSys)
	})
	r.Run(":8080")
}
