package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data := []byte("practicego")
	err := os.WriteFile("test.txt", data, 0222)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}