package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "¡Hola, edarcode!")
	})
	server.Logger.Fatal(server.Start(":3000"))
}
