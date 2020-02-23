package auth

import (
	"api/api"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"net/http"
)

type Handler struct {
}

func (h *Handler) Bind(e *echo.Echo) {
	e.GET("/oauth2/:provider", h.loginHandler())
	e.GET("/oauth2/:provider/callback", h.loginCallbackHandler())
}

func (h *Handler) loginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		provider := c.Param("provider")
		ctx := c.Request().Context()
		req := c.Request().WithContext(context.WithValue(ctx, "provider", provider))
		url, err := gothic.GetAuthURL(c.Response().Writer, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, api.Response{Message: err.Error()})
		}

		return c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func (h *Handler) loginCallbackHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		provider := c.Param("provider")
		ctx := c.Request().Context()
		req := c.Request().WithContext(context.WithValue(ctx, "provider", provider))
		u, err := gothic.CompleteUserAuth(c.Response().Writer, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, api.Response{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, api.Response{Success: true, Data: u})
	}
}

func NewHandler(providers ...goth.Provider) *Handler {
	goth.UseProviders(providers...)
	return &Handler{}
}
