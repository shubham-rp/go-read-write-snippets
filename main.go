package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// data := []byte("practicego")
	// err := os.WriteFile("test.txt", data, 0777)
	// if err !=nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("success")

	by,err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(by))
}