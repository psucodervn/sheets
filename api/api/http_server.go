package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/psucodervn/go/server"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	e := server.NewDefaultEchoServer()
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{Success: true, Message: "OK"})
	})

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
