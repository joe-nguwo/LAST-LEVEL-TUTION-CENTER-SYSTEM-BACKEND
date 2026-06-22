package handlers

import (
	"fmt"
	"last-level/database"
	"last-level/models"
	"net/http"

	_ "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := models.Admin{}

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "database error",
		})
	}

	result := db.Where("email = ? AND password = ?", email, password).First(&user)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid credentials",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "login successful",
	})
}

func Register(c echo.Context) error {
	fname := c.FormValue("firstName")
	lname := c.FormValue("lastName") // fixed typo
	password := c.FormValue("password")
	email := c.FormValue("email")

	register := models.Admin{
		Password: password,
		FName:    fname,
		LName:    lname,
		Email:    email,
	}

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("failed to connect to the database")

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "database error",
		
		})
	}

	result := db.Create(&register)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed to create user",
			
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "created user",
			
})
}
func Logout(c echo.Context) error {
	return c.JSON(http.StatusAccepted, map[string]string{"message": "logout "})
}

