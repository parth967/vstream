package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vstream/internal/handlers"
	"github.com/vstream/internal/pages"
)

func SetRouters(app *fiber.App) error {
	app.Get("/", func(c *fiber.Ctx) error {
		successMessage := c.Query("success")
		errorMessage := c.Query("error")

		return c.Render("layouts/index", fiber.Map{
			"SuccessMessage": successMessage,
			"ErrorMessage":   errorMessage,
		})
	}).Name("Index")

	app.Post("/login", handlers.HandleLogin).Name("Login")
	app.Post("/signup", handlers.HandleSignup).Name("Signup")

	app.Get("/d", DestroyServer).Name("ServerDestroy") ///remove this line and function later

	app.Use(handlers.AuthMiddleware)
	app.Get("/home", pages.RenderHome).Name("Home")
	app.Get("/settings", pages.RenderSettings).Name("Settings")

	return nil
}

func DestroyServer(ctx *fiber.Ctx) error {
	log.Fatal("Server Destroy")
	return nil
}
