package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookabledate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

// curl "localhost:8085/bookable?check_in=2030-04-16&check_out=2030-04-17"
// curl "localhost:8085/bookable?check_in=2030-03-10&check_out=2030-03-09"
// curl "localhost:8085/bookable?check_in=2000-03-09&check_out=2000-03-10"

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookabledate)
	}
	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBindWith(&b, binding.Query); err == nil {
			c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	r.Run(":8080")
}
