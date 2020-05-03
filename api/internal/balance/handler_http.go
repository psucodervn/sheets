package balance

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"api/internal/api"
)

type Handler struct {
	svc Service
}

func (h *Handler) Bind(e *echo.Echo) {
	e.GET("/users", h.getUsers(), api.QueryParser())
	e.GET("/users/:id", h.getUser())

	g := e.Group("/balance", api.QueryParser())
	g.GET("/transactions", h.getTransactions(), api.QueryParser())
	g.GET("/transactions/:id", h.getTransaction())
}

func (h *Handler) getUsersOld(c echo.Context) error {
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
		user, err := h.svc.User(ctx, userID)

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

func (h *Handler) getTransactionsOld() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		l := log.Ctx(ctx)

		txs, err := h.svc.FindTransactions(ctx, nil)
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

func (h *Handler) getUsers() echo.HandlerFunc {
	type request struct {
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
		args := api.QueryFromContext(c)
		users, err := h.svc.Users(ctx, args)
		if err != nil {
			l.Err(err).Msg("list users failed")
			return err
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: users})
	}
}

func (h *Handler) getTransactions() echo.HandlerFunc {
	type request struct {
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
		args := api.QueryFromContext(c)
		txs, err := h.svc.Transactions(ctx, args)
		if err != nil {
			l.Err(err).Msg("list transactions failed")
			return err
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: txs})
	}
}

func NewHandler(uc Service) *Handler {
	return &Handler{svc: uc}
}
