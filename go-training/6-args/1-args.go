package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[1:])
	fmt.Printf("%T", os.Args)
}
