package pages

import (
	"github.com/gofiber/fiber/v2"
)

func RenderSettingsPage(ctx *fiber.Ctx) error {
	settingData := fiber.Map{}
	return ctx.Render("layouts/settings", settingData)
}
