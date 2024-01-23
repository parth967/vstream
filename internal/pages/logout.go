package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vstream/internal/handlers"
)

func HandleLogout(ctx *fiber.Ctx) error {
	handlers.ClearUser(ctx)
	return ctx.Redirect("/")
}
