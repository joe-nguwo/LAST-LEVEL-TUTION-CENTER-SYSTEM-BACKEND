package handlers

import (
	db "last-level/database"
	"net/http"
	m "last-level/models"

	"github.com/labstack/echo/v4"
)

func StudentsStats(c echo.Context) error{
	
	db,err := db.ConnectDB()
	if err != nil{
		return  c.JSON(http.StatusInternalServerError,map[string]string{
			"message":"internal server error",
		})
	} 
	students := m.Students{}

	result := db.Find(&students)

	if result.Error != nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"message":"internal server error"},)

	}

	return c.JSON(http.StatusOK,map[string]string{
			"message":"internal server error",})

}

