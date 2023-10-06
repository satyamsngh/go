package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data := os.Args[1:]
	if len(data) != 3 {
		log.Println("provide all the values")
		return

	}
	fmt.Println("hello satyam")

}
