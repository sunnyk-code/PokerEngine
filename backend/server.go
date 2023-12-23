package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	e.GET("/", helloWorld)

	e.Start(":8080")
}
