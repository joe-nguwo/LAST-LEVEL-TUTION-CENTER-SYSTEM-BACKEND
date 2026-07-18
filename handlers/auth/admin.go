package handlers

import (
	"fmt"
	"last-level/database"
	"last-level/models"
	"last-level/types"
	"net/http"

	_ "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)


func Login(c echo.Context) error {
	var req types.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
	}

	user := models.Admin{}

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "database error",
		})
	}

	result := db.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid credentials",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid credentials",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "login successful",
	})
}

func Register(c echo.Context) error {
	var req types.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "error hashing password",
		})
	}

	register := models.Admin{
		FName:    req.FirstName,
		LName:    req.LastName,
		Email:    req.Email,
		Password: string(hashedPassword),
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
	return c.JSON(http.StatusAccepted, map[string]string{
		"message": "logout",
	})
}