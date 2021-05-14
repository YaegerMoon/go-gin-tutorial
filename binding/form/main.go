package main

import "github.com/gin-gonic/gin"

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var fakeForm myForm
		c.ShouldBind(&fakeForm)
		c.JSON(200, gin.H{"color": fakeForm.Colors})
	})

	r.Run(":8080")
}

// <form action="/" method="POST">
//     <p>Check some colors</p>
//     <label for="red">Red</label>
//     <input type="checkbox" name="colors[]" value="red" id="red">
//     <label for="green">Green</label>
//     <input type="checkbox" name="colors[]" value="green" id="green">
//     <label for="blue">Blue</label>
//     <input type="checkbox" name="colors[]" value="blue" id="blue">
//     <input type="submit">
// </form>
