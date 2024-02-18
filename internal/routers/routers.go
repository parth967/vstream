package routers

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/vstream/internal/handlers"
	"github.com/vstream/internal/pages"
)

func SetRouters(app *fiber.App) error {
	app.Get("/", pages.LoadHomePage).Name("Index")

	app.Post("/login", handlers.HandleLogin).Name("Login")
	app.Post("/signup", handlers.HandleSignup).Name("Signup")

	app.Get("/d", DestroyServer).Name("ServerDestroy") ///remove this line and function later
	app.Get("/r", RestartServer).Name("RestartServer") ///remove this line and function later

	app.Use(handlers.AuthMiddleware)
	app.Get("/home", pages.RenderHome).Name("Home")
	app.Get("/settings", pages.RenderSettingsPage).Name("Settings")
	app.Get("/getName", handlers.PrintName).Name("Name")
	app.Get("/logout", pages.HandleLogout).Name("Logout")
	app.Get("/getFilesMessage", pages.GetFilesMessage)
	app.Get("/folders", pages.FolderLists)
	app.Post("/processFolders", pages.ProcessFolders)

	return nil
}

func DestroyServer(ctx *fiber.Ctx) error {
	log.Fatal("Server Destroy")
	return nil
}

func RestartServer(ctx *fiber.Ctx) error {
	restartFile := "restart.sh"
	cmd := exec.Command("sh", restartFile)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing script:", err)
		return nil
	}

	// Print the output of the command
	fmt.Println("Script output:", string(output))

	return nil
}
