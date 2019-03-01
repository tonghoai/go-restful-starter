package main

import (
	"net/http"
	"restful-starter/db"
	"restful-starter/router"
	response "restful-starter/utilities"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	db.Ping()
}

func main() {
	server := router.Init()
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))
	server.Use(middleware.Recover())
	server.Use(middleware.BodyLimit("100K"))
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	server.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	server.Use(middleware.Static("static"))

	server.HTTPErrorHandler = customHTTPErrorHandler
	server.HideBanner = true
	server.Logger.Fatal(server.Start(":1234"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"
	if re, ok := err.(*echo.HTTPError); ok {
		code = re.Code
		msg = re.Message.(string)
	}
	c.JSON(code, response.KO(code, msg))
}
