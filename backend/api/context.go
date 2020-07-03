package api

import (
	"github.com/labstack/echo"
	"strings"
)

type eosbApiContext struct {
	echo.Context
	// If address is not empty means this user is authenticated.
	Address string
}

func initeosbApiContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &eosbApiContext{c, ""}
		return next(cc)
	}
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*eosbApiContext)
		cc.Response().Header().Set(echo.HeaderServer, "Echo/3.0")

		eosbAuthToken := cc.Request().Header.Get("Eosb-Authentication")
		eosbAuthTokens := strings.Split(eosbAuthToken, "#")

		if len(eosbAuthTokens) != 3 {
			return &ApiError{Code: -11, Desc: "Eosb-Authentication should be like {address}#EOSB-AUTHENTICATION@{time}#{signature}"}
		}

		valid, err := eosb.IsValidSignature(eosbAuthTokens[0], eosbAuthTokens[1], eosbAuthTokens[2])
		if !valid || err != nil {
			return &ApiError{Code: -11, Desc: "Eosb-Authentication valid failed, please check your authentication"}
		}
		cc.Address = strings.ToLower(eosbAuthTokens[0])
		return next(cc)
	}
}
