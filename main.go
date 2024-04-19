package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	app := echo.New()
    logger := app.Logger

    logger.SetLevel(log.DEBUG)

    app.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World")
    })
    app.Logger.Fatal(app.Start(":8080"))
}
