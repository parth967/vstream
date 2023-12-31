package servers

import (
	"errors"
	"fmt"
	"net/http"
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
	// TODO: Chaneg later to read proper log file and get info instead of constants
	// TODO: Add required info later on
	s.port = PORT
}

func Run() error {
	var sData serverInfo
	sData.initServer()
	sData.printServerInfo()

	if sData.port == "" {
		return errors.New("Port is not mentioned")
	}

	defaultPage := func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Default Page of vstream"))
	}
	http.HandleFunc("/", defaultPage)

	http.ListenAndServe(sData.port, nil)

	return nil
}
