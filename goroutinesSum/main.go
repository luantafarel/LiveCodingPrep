package main

import "fmt"

func sum(sl []int, c chan int) {
	cont := 0
	for _, val := range sl {
		cont += val
	}
	c <- cont
}
func main() {
	a := []int{1, 2, 3, 4, 5, 12, 3, 1, 2, 3, 4, 1, 23, 21}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x + y)
}
