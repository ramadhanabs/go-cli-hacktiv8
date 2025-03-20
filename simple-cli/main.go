package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	shout := flag.Bool("shout", false, "Print the message in uppercase")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: go run main.go <name> [--shout]")
		os.Exit(1)
	}

	name := args[0]
	message := fmt.Sprintf("Hello, %s!", name)

	if *shout {
		message = strings.ToUpper(message)
	}

	fmt.Println(message)
}
