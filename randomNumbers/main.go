package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("4 random numbers between 60 and 72:")

	for i := 0; i < 4; i++ {
		number := rand.Intn(13) + 60
		fmt.Printf("Number %d: %d\n", i+1, number)
	}
}
