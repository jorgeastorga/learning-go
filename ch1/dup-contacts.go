package main

//File that will read a list of contacts and find any duplicated

import (
	"bufio"
	"os"
)

func main() {
	//read the file from the command line arguments
	fileName := os.Args[1]

	//loop over each line of the text files
	file, _ := os.Open(fileName)

	//print each line
	input := bufio.NewScanner(file)

	for input.Scan() {
		println(input.Text())
	}
}
