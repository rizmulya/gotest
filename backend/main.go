package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	// server side
	app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "hi welcome"})
    })

    // frontend
    frontendRoutes := map[string]bool{
        "/about": true,
        "/contact": true,
        "/service": true,
    }

    app.Use(func(c *fiber.Ctx) error {
        if _, exists := frontendRoutes[c.Path()]; exists {
            return c.SendFile("./frontend/dist/index.html")
        }
        return c.Next()
    })

    app.Static("/", "./frontend/dist")

    app.Listen(":8000")
}
