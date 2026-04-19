package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error writing your file")
		panic(err)
	}

	output := Processor(string(file))

	err = os.WriteFile("output.txt", []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing your file")
		panic(err)
	}
	fmt.Println("File written successfully")
}
