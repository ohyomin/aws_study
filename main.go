package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Route

	e.GET("/hello", hello)

	e.GET("/ping", ping)

	e.Logger.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!")
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Pong")
}
