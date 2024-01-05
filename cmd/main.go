package main

import (
	"fmt"

	"gihub.com/vstream/internal/servers"
)

func main() {
	fmt.Println("Starting Server")
	servers.RunServer()
	fmt.Println("Server Stop")
}
