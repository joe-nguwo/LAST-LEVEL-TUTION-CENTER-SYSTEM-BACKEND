package routes

import (
	"github.com/labstack/echo/v4"
	"last-level/handlers/auth"

)

func AdminRoutes(e *echo.Echo) {

api := e.Group("/api/v1/admin")
api.POST("/auth/login",handlers.AdminHandler)
// api.POST("/auth/logout")
// api.POST("/auth/register")

}