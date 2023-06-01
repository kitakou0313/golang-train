package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world")
}

func main() {
	e := echo.New()

	e.GET("/", helloWorldHandler)

	e.Logger.Fatal(e.Start(":8081"))

}
