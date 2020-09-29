package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"api/internal/api"
	"api/model"
)

type DaysOffService interface {
	List(ctx context.Context, pager model.Pager) ([]model.DayOff, error)
	Create(ctx context.Context, req model.DayOffToCreate) (*model.DayOff, error)
	Delete(ctx context.Context, id string) error
}

type DaysOffHandler struct {
	svc    DaysOffService
	authMW echo.MiddlewareFunc
}

func NewDaysOffHandler(svc DaysOffService, authMW echo.MiddlewareFunc) *DaysOffHandler {
	return &DaysOffHandler{svc: svc, authMW: authMW}
}

func (h *DaysOffHandler) Bind(e *echo.Echo) {
	g := e.Group("/days-off")
	g.GET("", h.List)
	g.POST("", h.Create)
	g.DELETE("/:id", h.Delete)
}

func (h *DaysOffHandler) List(ec echo.Context) error {
	c := ec.(*api.Context)
	days, err := h.svc.List(c.Ctx(), c.Pager())
	if err != nil {
		return c.Err(http.StatusInternalServerError, err.Error())
	}
	return c.OK(days)
}

func (h *DaysOffHandler) Create(ec echo.Context) error {
	c := ec.(*api.Context)
	var req model.DayOffToCreate
	if err := c.Map(&req); err != nil {
		return err
	}
	do, err := h.svc.Create(c.Ctx(), req)
	if err != nil {
		log.Ctx(c.Ctx()).Err(err).Interface("req", req).Msg("Create failed")
		return c.Err(http.StatusInternalServerError, err.Error())
	}
	return c.OK(do)
}

func (h *DaysOffHandler) Delete(ec echo.Context) error {
	c := ec.(*api.Context)
	id := c.Param("id")
	if err := h.svc.Delete(c.Ctx(), id); err != nil {
		return err
	}
	return c.OK(nil)
}
