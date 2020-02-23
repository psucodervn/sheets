package balance

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Handler struct {
	uc Service
}

func (h *Handler) Bind(e *echo.Echo) {
	e.GET("/users", h.getUsers)
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

func NewHandler(uc Service) *Handler {
	return &Handler{uc: uc}
}
