package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func AdminHandler (c echo.Context) error {
	// email := c.FormValue("email")
	// name := c.FormValue("name")

	// if email == "nguwojoe@gmail.com"  && name == "Joe Nguwo" {
	// 	return c.JSON(http.StatusOK, map[string]string{"message": "Admin login successful", "id": "1"})
	// }

    return c.JSON(http.StatusOK, map[string]string{"message": "Admin login successful"})

}

func  Register(c echo.Context) error {
	return c.JSON(http.StatusOK,map[string]string{"message":"Register"})


}

func Logout(c echo.Context) error {
	return c.JSON(http.StatusAccepted,map[string]string{"message":"logout "})
}


