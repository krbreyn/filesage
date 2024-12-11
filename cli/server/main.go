package main

import (
	"fmt"
	"os"

	"github.com/krbreyn/filesage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR: expected command")
		return
	}

	switch os.Args[1] {
	case "start":
		filesage.StartServer()
	default:
		fmt.Println("ERROR: invalid command")
	}
}
