package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"

	"api/internal/api"
	"api/internal/user"
	"api/model"
)

type UserService interface {
	FindByAuthProvider(ctx context.Context, provider string, id string) (*model.User, error)
	AddAuthProvider(ctx context.Context, provider string, id string) (*model.AuthIdentity, error)
}

type TelegramService interface {
	GenerateLink(ctx context.Context, userID string) (string, error)
}

type Handler struct {
	authSvc     *Service
	userSvc     UserService
	telegramSvc TelegramService
	authMW      echo.MiddlewareFunc
}

func (h *Handler) Bind(e *echo.Echo) {
	e.POST("/auth/google", h.loginGoogle())
	e.GET("/auth/me", h.getMe(), h.authMW)
	e.POST("/auth/telegram", h.generateTelegramToken(), h.authMW)
}

func (h *Handler) loginGoogle() echo.HandlerFunc {
	type request struct {
		Code        string `json:"code" validate:"required"`
		RedirectURI string `json:"redirectUri" validate:"required"`
	}

	return func(ec echo.Context) error {
		c := ec.(*api.Context)
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}

		l := log.Ctx(c.Ctx())
		gu, err := h.authSvc.FetchGoogleUserWithCode(c.Ctx(), req.Code)
		if err != nil {
			l.Err(err).Msg("FetchGoogleUserWithCode failed")
			return c.Err(http.StatusBadRequest, "Invalid authorization code!")
		}

		u, err := h.userSvc.FindByAuthProvider(c.Ctx(), "google", gu.Email)
		if err != nil {
			l.Err(err).Str("email", gu.Email).Msg("FindByAuthProvider failed")
			if err != user.ErrAuthNotFound {
				return err
			}
			_, err = h.userSvc.AddAuthProvider(c.Ctx(), "google", gu.Email)
			if err != nil {
				l.Err(err).Str("email", gu.Email).Msg("AddAuthProvider failed")
				return err
			}
		}

		if u == nil {
			return c.Err(http.StatusBadRequest, "Your account was not activated. Please contact admin for support!")
		}

		u.Email = null.StringFrom(gu.Email) // TODO: get from db
		t, err := h.authSvc.SignWithUser(u)
		if err != nil {
			return err
		}

		return c.OK(Token{AccessToken: t})
	}
}

func (h *Handler) getMe() echo.HandlerFunc {
	return func(ec echo.Context) error {
		c := ec.(*api.Context)
		u := new(User).FromModel(c.User())
		return c.OK(u)
	}
}

func (h *Handler) generateTelegramToken() echo.HandlerFunc {
	return func(ec echo.Context) error {
		c := ec.(*api.Context)
		if !c.User().TelegramID.IsZero() {
			return c.Err(http.StatusBadRequest, "Your account was already integrated with Telegram.")
		}

		l := log.Ctx(c.Ctx())
		link, err := h.telegramSvc.GenerateLink(c.Ctx(), c.User().ID)
		if err != nil {
			l.Err(err).Msg("GenerateLink failed")
			return err
		}

		return c.OK(link)
	}
}

func NewHandler(authSvc *Service, userSvc UserService, telegramSvc TelegramService, authMW echo.MiddlewareFunc) *Handler {
	return &Handler{authSvc: authSvc, userSvc: userSvc, telegramSvc: telegramSvc, authMW: authMW}
}
