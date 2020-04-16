package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	message := "Hello, there!"
	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println(message)

	}
}
