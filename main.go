package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// 1. Append - Read & write 
	by,err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(by))

	msg := []byte("\nappended message")

	by = append(by, msg...)
	
	err = os.WriteFile("test.txt", by, 0777)
	if err !=nil {
		log.Fatal(err)
	}

	// 2. Write Operation
	// data := []byte("practicego")
	// err := os.WriteFile("test.txt", data, 0777)
	// if err !=nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("success")

	// 3. Read Operation
	// by,err := os.ReadFile("test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(by))
}