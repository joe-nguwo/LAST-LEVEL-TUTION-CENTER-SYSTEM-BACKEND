package main

import(
	"github.com/labstack/echo/v4"
	"log"
)


func main(){

	e :=  echo.New()

	e.GET("/admin", func(c echo.Context) error {

		return c.String(200, "Admin endpoint")
	})


	log.Println(e.Start(":8080"))

}