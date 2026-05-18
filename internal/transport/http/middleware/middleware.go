package middleware

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	claims "github.com/vaniax17/go-urlShortener/core/jwt"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return &claims.Claims{}
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})
}
