package middlewares

import (
	"octagon/dtos"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from cookie
		tokenString := c.Cookies("token")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing JWT Token",
			})
		}

		// Parse & Validate Token
		claims := &dtos.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return dtos.JWT_KEY, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid JWT Token",
			})
		}

		// Store email in the context locals
		c.Locals("user", claims.Email)

		return c.Next()
	}
}
