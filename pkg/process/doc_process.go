package process

import (
	"fmt"
	"log"
	"os"
)

func SetFolder(string path) error {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range files {
		fmt.Println(e)
	}
}
