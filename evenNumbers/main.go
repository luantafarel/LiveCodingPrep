package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("First 20 even numbers using 2 goroutines:")

	ch := make(chan int, 20)

	go func() {
		for i := 0; i < 40; i += 2 {
			ch <- i
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func() {
		for i := 100; i < 140; i += 2 {
			ch <- i
			time.Sleep(1 * time.Millisecond)
		}
	}()

	for i := 0; i < 20; i++ {
		num := <-ch
		fmt.Printf("%d° even number: %d\n", i+1, num)
	}
}
