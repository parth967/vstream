package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetName(ctx *fiber.Ctx) string {
	data := ctx.Locals("name")
	if data != nil {
		return data.(string)
	}
	return ""
}

func PrintName(ctx *fiber.Ctx) error {
	return ctx.SendString(GetName(ctx))
}
