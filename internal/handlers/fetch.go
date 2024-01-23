package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetName(ctx *fiber.Ctx) string {
	return ctx.Locals("name").(string)
}

func PrintName(ctx *fiber.Ctx) error {
	return ctx.SendString(GetName(ctx))
}
