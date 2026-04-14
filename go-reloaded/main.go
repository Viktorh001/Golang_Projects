package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error running your code")
		panic(err)
	}
	input := string(data)

	err = os.WriteFile("output.txt", []byte{input}, 0644)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Written successfully")
}
