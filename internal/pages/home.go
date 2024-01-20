package pages

import (
	"github.com/gofiber/fiber/v2"
)

func RenderHome(ctx *fiber.Ctx) error {
	Name, usernameOk := ctx.Locals("username").(string)

	if !usernameOk {
		ctx.Redirect("/")
	} else {
		data := fiber.Map{
			"Title": "VSTREAM - Home",
			"Name":  Name,
		}

		return ctx.Render("layouts/home", data)
	}
	return nil
}
