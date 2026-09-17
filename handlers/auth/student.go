package handlers

import (
	"fmt"
	"net/http"

	"last-level/database"
	"last-level/models"
	"last-level/types"

	"github.com/labstack/echo/v4"
)

func RegisterStudent(c echo.Context) error {
	var req types.RegisterStudentRequest

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

func AllStudents(c echo.Context) error{
	db,err := database.ConnectDB()


	if err != nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"message":"cant connect to the database",
		})
	}
	

	var students []models.Students
	if err := db.Find(&students).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to fetch students",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":"ok",
		"data": students,
	})

}

func DeleteStudent(c echo.Context) error {
	db,err := database.ConnectDB()
	users := models.Students{}
	var req types.DeleteStudent
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
	}

	if err != nil{
		return  c.JSON(http.StatusInternalServerError,map[string]string{
			"message":"database connection failed",
		})

	}
	id := c.Param("id")
	fmt.Println("this is the id",id)
	results := db.Delete(&users,id)
	if(results.RowsAffected > 0){
		fmt.Println("delete succefull")
	}else{
			fmt.Println(results.Error,"error occures while deleting")

	}
	


return c.NoContent(http.StatusNoContent)
	
	

}

