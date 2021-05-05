package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -X POST http://localhost:8080/upload \
//   -F "upload[]=@/Users/appleboy/test1.zip" \
//   -F "upload[]=@/Users/appleboy/test2.zip" \
//   -H "Content-Type: multipart/form-data"

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 24
	router.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, fmt.Sprintf("./%s", file.Filename))
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
	})

	router.Run(":8080")
}
