package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	// Function that will cause a panic
	causePanic()
}

func causePanic() {
	file, err := os.Open("nonexistent_file.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer file.Close()
}
