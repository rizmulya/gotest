package main

import (
	"gotest/database"
	"gotest/handlers"
	"gotest/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
	/*
	 * setup
	 */
	database.Connect()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// for development only
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	/*
	 * Server Side API
	 */
	// /
	app.Post("/logout", middleware.Auth(), handlers.Logout)
	// /api
	api := app.Group("/api")
	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)
	// /api/users
	user := api.Group("/users", middleware.Auth("admin"))
	user.Get("/", handlers.GetUsers)
	user.Get("/:uid", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Put("/:uid", handlers.UpdateUser)
	user.Delete("/:uid", handlers.DeleteUser)

	/*
	 * Server Side Page
	 */
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("landingPage", fiber.Map{
			"key": "value",
		})
	})

	/*
	 * Client Side Page
	 */
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"key": "valuex",
		})
	})
	app.Get("/users", middleware.Auth(), func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"key": "valuex",
		})
	})

	// app.Static("/", "./frontend/dist")

	app.Listen(":8000")
}
