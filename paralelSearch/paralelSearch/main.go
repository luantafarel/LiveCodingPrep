package main

import "fmt"

func search(n []int, s int, c chan bool) {
	c <- false
	for _, val := range n {
		if val == s {
			c <- true
			return
		}
	}
	return
}

func main() {
	c := make(chan bool)
	arr := []int{1, 2, 3, 1, 12, 3, 1, 2, 3, 191, 1, 22, 1838, 10398, 11}
	go search(arr[len(arr)/2:], 12111, c)
	go search(arr[:len(arr)/2], 12111, c)
	x, y := <-c, <-c
	fmt.Println(x || y)
	return
}
