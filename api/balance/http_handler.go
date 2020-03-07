package balance

import (
	"api/api"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Handler struct {
	uc Service
}

func (h *Handler) Bind(e *echo.Echo) {
	e.GET("/users", h.getUsers)
	e.GET("/users/:id", h.getUser())
}

func (h *Handler) getUsers(c echo.Context) error {
	ctx := c.Request().Context()
	// users, err := h.uc.ListUserBalances(ctx)
	users, err := h.uc.Find(ctx, nil)

	l := log.Ctx(ctx)
	if err != nil {
		l.Err(err).Msg("fetch users failed")
		c.Error(err)
		return nil
	}

	return c.JSON(http.StatusOK, api.Response{Success: true, Data: users})
}

func (h *Handler) getUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userID := c.Param("id")
		user, err := h.uc.FindByID(ctx, userID)

		l := log.Ctx(ctx)
		if errors.Is(err, ErrNotFound) {
			return c.JSON(http.StatusNotFound, api.Response{Message: "User not found"})
		}
		if err != nil {
			l.Err(err).Str("userID", userID).Msg("FindByID failed")
			return c.JSON(http.StatusInternalServerError, api.Response{})
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: user})
	}
}

func NewHandler(uc Service) *Handler {
	return &Handler{uc: uc}
}
