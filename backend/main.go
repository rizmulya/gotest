package main

import (
	"gotest/database"
	"gotest/handlers"
	"gotest/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"time"
)

func main() {
	/*
	 * setup
	 */
	database.Connect()

	store := session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   false, //dev
		CookieSameSite: "Strict",
		Expiration:     24 * time.Hour,
	})

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// for development only
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, X-CSRF-Token",
		AllowCredentials: true,
	}))

	// CSRF middleware
	app.Use(csrf.New(csrf.Config{
		KeyLookup:         "header:X-Csrf-Token",
		CookieName:        "csrf_",
		CookieSameSite:    "Strict",
		CookieSecure:      false, //dev
		CookieHTTPOnly:    false,
		CookieSessionOnly: true,
		Expiration:        1 * time.Hour,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid CSRF token",
			})
		},
		Session:    store,
		SessionKey: "fiber.csrf.token",
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
		return c.Render("landingPage", fiber.Map{})
	})

	/*
	 * Client Side Page
	 */
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	app.Get("/users", middleware.Auth(), func(c *fiber.Ctx) error {
		return c.SendFile("./views/index.html")
	})

	app.Listen(":8000")
}