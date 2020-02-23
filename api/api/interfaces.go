package api

import "github.com/labstack/echo/v4"

type EchoHandler interface {
	Bind(e *echo.Echo)
}
