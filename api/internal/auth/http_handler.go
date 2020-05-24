package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"

	"api/internal/api"
	"api/model"
)

type UserService interface {
	FindByAuthProvider(ctx context.Context, provider string, id string) (*model.User, error)
}

type Handler struct {
	authSvc *Service
	userSvc UserService
	authMW  echo.MiddlewareFunc
}

func (h *Handler) Bind(e *echo.Echo) {
	e.POST("/auth/google", h.loginGoogle())
	e.GET("/auth/me", h.getMe(), h.authMW)
}

func (h *Handler) loginGoogle() echo.HandlerFunc {
	type request struct {
		Code        string `json:"code" validate:"required"`
		RedirectURI string `json:"redirectUri" validate:"required"`
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
		gu, err := h.authSvc.FetchGoogleUserWithCode(ctx, req.Code)
		if err != nil {
			l.Err(err).Msg("FetchGoogleUserWithCode failed")
			return err
		}

		u, err := h.userSvc.FindByAuthProvider(ctx, "google", gu.Email)
		if err != nil {
			l.Err(err).Str("email", gu.Email).Msg("FindByAuthProvider failed")
			return err
		}

		u.Email = null.StringFrom(gu.Email) // TODO: get from db
		t, err := h.authSvc.SignWithUser(u)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, api.Response{Success: true, Data: echo.Map{
			"accessToken": t,
		}})
	}
}

func (h *Handler) getMe() echo.HandlerFunc {
	return func(ec echo.Context) error {
		c := ec.(*api.Context)
		u := new(User).FromModel(c.User())
		return c.OK(u)
	}
}

func NewHandler(authSvc *Service, userSvc UserService, authMW echo.MiddlewareFunc) *Handler {
	return &Handler{authSvc: authSvc, userSvc: userSvc, authMW: authMW}
}
