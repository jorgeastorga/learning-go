package main

import (
	"fmt"
	"os"
)

func main() {
	var numArguments = len(os.Args)
	fmt.Println(numArguments)
	fmt.Println(os.Args[0])
}
