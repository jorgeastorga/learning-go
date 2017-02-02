package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the CLI - Get your contacts")
	fmt.Println(os.Args[0])
	var cliArguments = os.Args[1:]
	fmt.Println(cliArguments)

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("Jorge")
	}

	fmt.Println(s)
}
