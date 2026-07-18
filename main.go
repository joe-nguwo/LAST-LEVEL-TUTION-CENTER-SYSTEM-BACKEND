package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"last-level/database"
	_ "last-level/handlers/auth"
	"last-level/routes"

	"log"
	//mid "last-level/middleware"
)

func main() {

	e := echo.New()

	_ = godotenv.Load()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	origins := []string{
		"http://localhost:3000",
		"http://localhost:5173",
		"http://localhost:5174",
		"http://localhost:8081",
	}
	database.MigrateDb(db)
	
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     origins,
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS, echo.PATCH},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours for preflight cache
	}))
	e.Use(middleware.Logger())

	//e.Use(mid.RequestTime) 

	routes.AdminRoutes(e)

	log.Println(e.Start(":8000"))

}
