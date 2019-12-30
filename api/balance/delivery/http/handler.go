package http

import (
	"api/balance"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Handler struct {
	uc balance.UseCase
}

func (h *Handler) getUsers(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := h.uc.ListUserBalances(ctx)

	l := log.Ctx(c.Request().Context())
	if err != nil {
		l.Err(err).Msg("fetch users failed")
		c.Error(err)
		return nil
	}

	return c.JSON(http.StatusOK, users)
}

func NewHandler(uc balance.UseCase) *Handler {
	return &Handler{uc: uc}
}

func BindEchoHandler(e *echo.Echo, uc balance.UseCase) {
	h := NewHandler(uc)
	e.GET("/users", h.getUsers)
}
