package domain

import (
	"github.com/labstack/echo/v4"
)

func NewContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Context{
			Context:   c,
			PubKey:    nil,
			SecretKey: nil,
		}
		return next(cc)
	}
}

type Context struct {
	echo.Context
	PubKey    *UserPubKey
	SecretKey *UserSecretKey
}
