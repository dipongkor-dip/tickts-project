package middlewares

import (
	"gotickts/internal/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

// ServerHeader middleware adds a `Server` header to the response.

func AuthMiddleware(jwtService auth.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			// extra token from authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing authorization header"})
			}
			// check bearer schema
			parts := strings.Split(authHeader, " ") // ["Bearer", "token"]

			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization header format"})
			}

			tokenString := parts[1]
			// validate token
			claims, err := jwtService.ValidateToken(tokenString)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
			}

			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)
			c.Set("user_name", claims.Name)
			// store user info in context for handlers
			return next(c)
		}
	}
}
