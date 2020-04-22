package api

import (
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/psucodervn/go/logger"
	"github.com/psucodervn/go/validator"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(logger.EchoMiddleware(func(c echo.Context) bool {
		uri := c.Request().RequestURI
		return uri == "/healthz" || uri == "/metrics"
	}))
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{Success: true, Message: "OK"})
	})
	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})
	e.Validator = validator.NewStructValidator()

	p := prometheus.NewPrometheus("http", nil)
	p.Use(e)

	srv := &Server{e: e}
	return srv
}

func (s *Server) Serve(addr string, tls bool) error {
	if tls {
		return s.e.StartAutoTLS(addr)
	}
	return s.e.Start(addr)
}

func (s *Server) Bind(handlers ...EchoHandler) {
	for i := range handlers {
		handlers[i].Bind(s.e)
	}
}
