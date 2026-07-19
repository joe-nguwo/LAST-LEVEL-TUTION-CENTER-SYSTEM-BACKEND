package handlers

import (
	"net/http"

	"last-level/database"
	"last-level/models"
	"last-level/types"

	"github.com/labstack/echo/v4"
)

func RegisterStudent(c echo.Context) error {
	var req types.RegisterStudentRequest

	// Bind JSON request to struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
	}

	db, err := database.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "database connection failed",
		})
	}

	student := models.Students{
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Email:      req.Email,
		Contact:    req.Contact,
		Address:    req.Address,
	}

	if err := db.Create(&student).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to register student",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "student registered successfully",
		"student": student,
	})
}