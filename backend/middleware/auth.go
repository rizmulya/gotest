package middleware

import (
    "gotest/utils"
    "github.com/gofiber/fiber/v2"
	"strings"
)

func Auth(roles ...string) fiber.Handler {
    return func(c *fiber.Ctx) error {
		// use token from header
        // authHeader := c.Get("Authorization")
        // tokenString := extractToken(authHeader)

		// use token from cookie
		tokenString := c.Cookies("jwt")

        if tokenString == "" {
            return c.Redirect("/login", fiber.StatusFound)
            // return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access Denied!"})
        }

        claims, err := utils.ParseJWT(tokenString)
        if err != nil {
            return c.Redirect("/login", fiber.StatusFound)
            // return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Invalid or expired JWT"})
        }

        c.Locals("userID", claims.UserID)
        c.Locals("userRole", claims.Role)

        if len(roles) > 0 {
            authorized := false
            for _, role := range roles {
                if claims.Role == role {
                    authorized = true
                    break
                }
            }
            if !authorized {
                return c.Redirect("/login", fiber.StatusFound)
                // return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized!"})
            }
        }

        return c.Next()
    }
}

func extractToken(authHeader string) string {
    if strings.HasPrefix(authHeader, "Bearer ") {
        return strings.TrimPrefix(authHeader, "Bearer ")
    }
    return ""
}