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

    app.Listen(":8000")
}
