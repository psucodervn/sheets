package point

import (
	"api/api"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

type HttpHandler struct {
	svc Service
}

func NewHttpHandler(svc Service) *HttpHandler {
	return &HttpHandler{svc: svc}
}

func (h *HttpHandler) Bind(e *echo.Echo) {
	e.GET("/points", h.getPoints())
}

func (h *HttpHandler) getPoints() echo.HandlerFunc {
	type request struct {
		Month int `json:"month" validate:"required,min=1,max=12"`
		Year  int `json:"year" validate:"required,min=2019,max=2020"`
	}

	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, api.Response{Message: err.Error()})
		}
		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, api.Response{Message: err.Error()})
		}

		ctx := c.Request().Context()
		l := log.Ctx(ctx).With().Interface("req", req).Logger()
		up, err := h.svc.UserPoints(ctx, req.Month, req.Year)
		if err != nil {
			l.Err(err).Msg("UserPoints failed")
			return c.JSON(http.StatusServiceUnavailable, api.Response{Message: "Service unavailable."})
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: up})
	}
}
