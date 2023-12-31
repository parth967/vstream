package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("rooms/%s", gguid.New().String))
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}

	uuid, suuid, _ := createORGetRoom(uuid)
}

func RoomWebSocket(c *websocket.Conn) {
	uuid := c.params("uuid")
	if uuid == "" {
		return
	}

	_, _, room := createORGetRoom(uuid)

}

func createORGetRoom(uuid string) (string, string, Room) {

}
