package point

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"api/internal/api"
)

type HttpHandler struct {
	svc       Service
	reportSvc ReportService
}

func NewHttpHandler(svc Service, reportSvc ReportService) *HttpHandler {
	return &HttpHandler{svc: svc, reportSvc: reportSvc}
}

func (h *HttpHandler) Bind(e *echo.Echo) {
	e.GET("/points", h.getPoints())
	e.GET("/report", h.getReport())
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
		up, err := h.svc.UserPointsInMonth(ctx, req.Month, req.Year)
		if err != nil {
			l.Err(err).Msg("UserPointsInMonth failed")
			return c.JSON(http.StatusServiceUnavailable, api.Response{Message: "Service unavailable."})
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: up})
	}
}

func (h *HttpHandler) getReport() echo.HandlerFunc {
	type request struct {
		From api.Timestamp `json:"from" validate:"required"`
		To   api.Timestamp `json:"to" validate:"required"`
	}

	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}

		ctx := c.Request().Context()
		l := log.Ctx(ctx)
		resp, err := h.reportSvc.GetReport(ctx, time.Time(req.From), time.Time(req.To))
		if err != nil {
			l.Err(err).Msg("GetReport failed")
			return err
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: resp})
	}
}
