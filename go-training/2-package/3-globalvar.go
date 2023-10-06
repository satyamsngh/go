package main

import (
	"fmt"
	"go-training/sum"
)

func main() {

	sum.Add()

	fmt.Println(sum.Sum)
	sum.Sum = 100
	fmt.Println(sum.Sum)
}
