package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("cookie_name")
		print(cookie, err)
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("cookie_name", "test", 300, "/", "localhost", false, false)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	r.Run(":8080")
}
