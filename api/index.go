package api

import (
	"net/http"
	"restful-starter/utilities"

	"github.com/labstack/echo"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, utilities.OK(map[string]interface{}{"message": "Golang API v1.0"}))
	}
}
