package http

import (
	"api/balance"
	balanceHttp "api/balance/delivery/http"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/psucodervn/go/logger"
	"net/http"
)

type Server struct {
	e *echo.Echo
}

func NewServer(balanceUseCase balance.UseCase) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(logger.EchoMiddleware(func(c echo.Context) bool {
		uri := c.Request().RequestURI
		return uri == "/healthz" || uri == "/metrics"
	}))
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"success": true, "message": "OK"})
	})

	p := prometheus.NewPrometheus("http", nil)
	p.Use(e)

	balanceHttp.BindEchoHandler(e, balanceUseCase)

	srv := &Server{e: e}
	return srv
}

func (s *Server) Serve(addr string, tls bool) error {
	if tls {
		return s.e.StartAutoTLS(addr)
	}
	return s.e.Start(addr)
}
