package main

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func HomeHandler(c echo.Context) error {
	c.Logger().Debug("got here")
	return Render(c, http.StatusOK, Home())
}

func TimeHandler(ctx echo.Context) error {
	//return TimeHandler{Now: now}
	ctx.Logger().Debug("got here")
	return Render(ctx, http.StatusOK, timeComponent(time.Now()))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func main() {
	app := echo.New()
	logger := app.Logger

	logger.SetLevel(log.DEBUG)

	app.GET("/", HomeHandler)
	app.GET("/time", TimeHandler)
	app.Logger.Fatal(app.Start(":8080"))
}
