package servers

import (
	"fmt"

	"time"

	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()
	now := time.Now().Format("2006-01-02 15:04:05")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Default Page of vStream ..... " + now)
	})
	app.Get("/auth")

	app.Listen(s.port)
}
