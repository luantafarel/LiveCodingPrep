package main

import "fmt"

func sum(n []int, c chan int) {
	sum := 0
	for _, val := range n {
		sum += val
	}
	c <- sum
}

func main() {
	c := make(chan int)
	ar := []int{1, 2, 3, 4, 1, 2, 3, 5, 10}
	go sum(ar[len(ar)/2:], c)
	go sum(ar[:len(ar)/2], c)
	x, y := <-c, <-c
	fmt.Println(x + y)
	return
}
