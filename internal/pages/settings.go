package pages

import (
	"github.com/gofiber/fiber/v2"
)

func RenderSettings(ctx *fiber.Ctx) error {
	return ctx.Redirect("/settingsPage")
}
