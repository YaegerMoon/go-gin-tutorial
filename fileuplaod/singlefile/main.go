package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -X POST http://localhost:8080/upload \
//   -F "file=@/Users/appleboy/test.zip" \
//   -H "Content-Type: multipart/form-data"

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 25
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		err := c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", "./", file.Filename))
		if err != nil {
			fmt.Print(err)
			c.String(http.StatusExpectationFailed, "file to uplaod")
		}
		c.String(http.StatusOK, fmt.Sprintf("%s uploaded!", file.Filename))
	})
	router.Run(":8080")
}
