package servers

import (
	"fmt"

	"time"

	"github.com/gofiber/fiber/v2"

	// To use a specific template engine, import as shown below:
	// "github.com/gofiber/template/pug"
	// "github.com/gofiber/template/mustache"
	// etc..

	// In this example we use the html template engine
	"github.com/gofiber/template/html/v2"
)

const PORT = ":8080"

type ServerData interface {
	getserverInfo()
	printServerInfo()
}

type serverInfo struct {
	port string
}

func (s *serverInfo) printServerInfo() {
	fmt.Printf("Listening port: %v", s.port)
}

func (s *serverInfo) initServer() {
	s.port = PORT
}

func RunServer() {
	var s serverInfo
	s.initServer()
	s.printServerInfo()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	now := time.Now().Format("2006-01-02 15:04:05")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("layouts/index", fiber.Map{
			"Title":       "vstream-app",
			"Description": "This is default Page of app" + now,
		})
	})

	app.Listen(s.port)
}
