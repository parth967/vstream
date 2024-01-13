package servers

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/vstream/internal/handlers"
)

type ServerData interface {
	getserverInfo()
}

type serverInfo struct {
	port            string
	serverStartTime string
}

func (s *serverInfo) initServer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading in .env file")
	}

	s.port = ":" + os.Getenv("PORT")
	s.serverStartTime = time.Now().Format("2006-01-02 15:04:05")
}

func setRouters(app *fiber.App) error {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("layouts/index", fiber.Map{})
	})
	app.Post("/login", handlers.HandleLogin)
	app.Post("/signup", handlers.HandleSignup)
	return nil
}

func RunServer() {
	var server serverInfo
	server.initServer()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "./assets")
	app.Use(logger.New()) //TODO: Modify logger later on and on off based on .env file

	setRouters(app)

	app.Listen(server.port)
}
