package main

import (
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
	"last-level/database"
	_"last-level/handlers/auth"
	"log"
	"last-level/routes"
	"last-level/middleware"

	
)

func main() {

	e := echo.New()

	_ = godotenv.Load()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	database.MigrateDb(db)
	e.Use(middleware.RequestTime)


	routes.AdminRoutes(e)

	log.Println(e.Start(":3000"))

}
