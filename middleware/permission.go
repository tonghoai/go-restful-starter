package middleware

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Permission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		users := c.Get("user").(*jwt.Token)
		claims := users.Claims.(jwt.MapClaims)
		if int(claims["permission"].(float64)) == 0 {
			return echo.NewHTTPError(http.StatusForbidden, "Bạn không có quyền")
		}
		next(c)
		return nil
	}
}
