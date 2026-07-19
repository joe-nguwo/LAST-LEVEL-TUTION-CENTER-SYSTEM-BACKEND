package routes

import (
	"github.com/labstack/echo/v4"
	"last-level/handlers/auth"

)

func AdminRoutes(e *echo.Echo) {

api := e.Group("/api/v1/admin")
api.POST("/auth/login",handlers.Login)
api.POST("/auth/logout",handlers.Logout)
api.POST("/auth/register",handlers.Register)
// Dashbaord Statitics 
api.GET("/auth/dashStats",handlers.StudentsStats)
// students
api.POST("/auth/student",handlers.RegisterStudent)
api.GET("/auth/student",handlers.AllStudents)
}