package main

import(
	"github.com/labstack/echo/v4"
	"log"
	"last-level/database"
	"last-level/handlers/auth"


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

	e.GET("/admin/login/:id",handlers.AdminHandler)
	e.GET("/Query",handlers.Show)


	log.Println(e.Start(":8080"))

}