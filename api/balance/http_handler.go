package balance

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"api/api"
)

type Handler struct {
	svc Service
}

func (h *Handler) Bind(e *echo.Echo) {
	e.GET("/users", h.getUsers)
	e.GET("/users/:id", h.getUser())

	g := e.Group("/balance", api.QueryParser())
	g.GET("/transactions", h.getTransactions())
	g.GET("/transactions/:id", h.getTransaction())
}

func (h *Handler) getUsers(c echo.Context) error {
	ctx := c.Request().Context()
	// users, err := h.svc.ListUserBalances(ctx)
	users, err := h.svc.FindUsers(ctx, nil)

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
		user, err := h.svc.FindUserByID(ctx, userID)

		l := log.Ctx(ctx)
		if errors.Is(err, ErrNotFound) {
			return c.JSON(http.StatusNotFound, api.Response{Message: "User not found"})
		}
		if err != nil {
			l.Err(err).Str("userID", userID).Msg("FindUserByID failed")
			return c.JSON(http.StatusInternalServerError, api.Response{})
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: user})
	}
}

func (h *Handler) getTransactions() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		l := log.Ctx(ctx)

		args := api.QueryFromContext(c)
		if len(args.OrderBy) == 0 {
			args.OrderBy = "time"
			args.Descending = true
		}
		if args.Limit <= 0 || args.Limit >= 1000 {
			args.Limit = 1000
		}
		txs, err := h.svc.FindTransactions(ctx, &args)
		if err != nil {
			l.Err(err).Msg("FindTransactions failed")
			return c.JSON(http.StatusInternalServerError, api.Response{})
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: txs})
	}
}

func (h *Handler) getTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		l := log.Ctx(ctx)

		txID := c.Param("id")
		tx, err := h.svc.FindTransaction(ctx, txID)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				return c.JSON(http.StatusNotFound, api.Response{})
			}
			l.Err(err).Str("txID", txID).Msg("FindTransaction failed")
			return c.JSON(http.StatusInternalServerError, api.Response{})
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: tx})
	}
}

func NewHandler(uc Service) *Handler {
	return &Handler{svc: uc}
}
