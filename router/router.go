package router

import (
	"restful-starter/api"
	permission "restful-starter/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	config := middleware.JWTConfig{
		SigningKey: []byte("secret"),
		ContextKey: "user",
	}
	e.GET("/", api.Index())
	e.POST("/login", api.Login())
	e.POST("/register", api.Register())

	user := e.Group("/user")
	user.Use(middleware.JWTWithConfig(config))
	user.Use(permission.Permission)
	user.GET("", api.GetUser())

	return e
}
