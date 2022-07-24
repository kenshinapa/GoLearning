package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No arguments provided. Provide a file name for the program to work")
		os.Exit(1)
	}
	fileLocation := os.Args[1]
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Error opening file: ", fileLocation, "Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
