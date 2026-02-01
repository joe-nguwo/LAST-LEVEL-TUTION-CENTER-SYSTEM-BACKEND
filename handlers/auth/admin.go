package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func AdminHandler (c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK," you id is " + id)

}

func Show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
