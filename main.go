package main

import(
	"github.com/labstack/echo/v4"
	"log"
	"last-level/database"

	"github.com/joho/godotenv"
)


func main(){

	e :=  echo.New()

	_ = godotenv.Load()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	database.MigrateDb(db)

	e.GET("/admin", func(c echo.Context) error {

		return c.String(200, "Admin endpoint")
	})


	log.Println(e.Start(":8080"))

}