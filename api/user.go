package api

import (
	"fmt"
	"net/http"
	"restful-starter/handle"
	"restful-starter/models"
	response "restful-starter/utilities"

	validator "gopkg.in/go-playground/validator.v9"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var validate *validator.Validate

func Register() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		validate = validator.New()
		u := new(models.User)
		if err = c.Bind(u); err != nil {
			return err
		}

		err = validate.Struct(u)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Dữ liệu không hợp lệ")
		}

		user, err := handle.Register(u)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Lỗi đăng ký")
		}
		return c.JSON(http.StatusOK, response.OK(map[string]interface{}{"user": user}))
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		validate = validator.New()
		u := new(models.User)
		if err = c.Bind(u); err != nil {
			return err
		}

		err = validate.Struct(u)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Dữ liệu không hợp lệ")
		}

		token, err := handle.Login(u)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Lỗi đăng nhập")
		}

		return c.JSON(http.StatusOK, response.OK(map[string]interface{}{"token": token}))
	}
}

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		users := c.Get("user").(*jwt.Token)
		claims := users.Claims.(jwt.MapClaims)
		fmt.Println(claims)

		user := models.User{ID: int(claims["id"].(float64))}
		u, err := handle.GetUser(&user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Có lỗi sảy ra")
		}
		return c.JSON(http.StatusOK, response.OK(map[string]interface{}{"user": u}))
	}
}
