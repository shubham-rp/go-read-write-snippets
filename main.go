package main

import (
	"log"
	"os"
)

func main() {

	// 1-A. Append - Read & write 
	// by,err := os.ReadFile("test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(by))

	// msg := []byte("\nappended message")

	// by = append(by, msg...)
	
	// err = os.WriteFile("test.txt", by, 0777)
	// if err !=nil {
	// 	log.Fatal(err)
	// }

	// 1-B. Append - Open & Write 
	
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	msg := []byte("\nanother appended message")

	_,err = file.Write(msg)
	if err != nil {
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