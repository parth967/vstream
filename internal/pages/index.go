package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vstream/internal/handlers"
)

func LoadHomePage(ctx *fiber.Ctx) error {
	if ctx.Cookies("Auth") != "" {
		handlers.AuthMiddleware(ctx)
		name := handlers.GetName(ctx)
		if name != "" {
			ctx.Redirect("/home")
		}
	}
	return ctx.Render("layouts/index", fiber.Map{})
}
