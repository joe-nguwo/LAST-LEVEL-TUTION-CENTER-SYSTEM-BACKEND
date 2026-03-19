package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)


func RequestTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("the time of the request is ",time.Now())
		return next(c)
	}
}

