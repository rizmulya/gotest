package main

import (
    "gotest/database"
    "gotest/handlers"
    "gotest/middleware"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    /*
    * setup
    */
    database.Connect()

    engine := html.New("./views", ".tmpl")

    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // for development only
    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173",
        AllowHeaders: "Origin, Content-Type, Accept",
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
    user.Get("/:id", handlers.GetUser)
    user.Post("/", handlers.CreateUser)
    user.Put("/:id", handlers.UpdateUser)
    user.Delete("/:id", handlers.DeleteUser)


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
        return c.SendFile("./frontend/dist/index.html")
    })
    app.Get("/users", middleware.Auth(), func(c *fiber.Ctx) error {
        return c.SendFile("./frontend/dist/index.html")
    })

    app.Static("/", "./frontend/dist")

    app.Listen(":8000")
}
