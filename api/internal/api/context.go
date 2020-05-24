package api

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"api/model"
)

type Context struct {
	echo.Context
}

func (c *Context) User() *model.User {
	token, ok := c.Get(KeyUser).(*jwt.Token)
	if !ok {
		return nil
	}
	u, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil
	}

	user, _ := model.FindUserG(c.Request().Context(), u.ID)
	return user
}

func (c *Context) Ctx() context.Context {
	return c.Request().Context()
}

func (c *Context) OK(data interface{}, code ...int) error {
	co := http.StatusOK
	if len(code) > 0 {
		co = code[0]
	}
	return c.JSON(co, Response{Success: true, Data: data})
}

func (c *Context) Err(code int, msg string) error {
	return c.JSON(code, Response{Success: false, Message: msg})
}
