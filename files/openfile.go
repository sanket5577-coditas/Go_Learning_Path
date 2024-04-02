package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Go")

	content := "This needs to go in a file - MrSanketKumar"
	filename := "./myfile.txt"

	// Check if the file already exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// File doesn't exist, so create it
		file, err := os.Create(filename)
		checkNilErr(err)
		defer file.Close()

		// Write content to the file
		length, err := io.WriteString(file, content)
		checkNilErr(err)
		fmt.Println("Length is ", length)
	} else {
		fmt.Println("File already exists")
	}

	// Read the file
	readFile(filename)
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
