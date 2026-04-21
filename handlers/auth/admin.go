package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	//"github.com/golang-jwt/jwt/v5"
)




func AdminHandler (c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	

	if email  !="nguwojoe@gmail.com" || password != "1234"{
		
		return  c.JSON(http.StatusOK,map[string]string{"token":"1234464"})
		
	}else{
		    return c.JSON(http.StatusNotFound, map[string]string{"message": "error while loggin"})

	}

	
}

func  Register(c echo.Context) error {
	return c.JSON(http.StatusOK,map[string]string{"message":"Register"})


}

func Logout(c echo.Context) error {
	return c.JSON(http.StatusAccepted,map[string]string{"message":"logout "})
}


