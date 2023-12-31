package main

import (
	"fmt"

	"gihub.com/vstream/internal/servers"
)

func main() {
	fmt.Println("Starting Server")
	servers.Run()
	fmt.Println("Server Stop")
}
