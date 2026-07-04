package handlers

import (
	db "last-level/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func StudentsStats(c echo.Context) error {
	conn, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "database connection error",
		})
	}

	var total int64

	if err := conn.Table("students").Count(&total).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to count students",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"total_students": total,
	})
}