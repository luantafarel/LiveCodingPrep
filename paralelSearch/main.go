package main

import "fmt"

func search(sl []int, num int, ch chan bool) {
	ch <- false
	for i := 0; i < len(sl); i++ {
		if sl[i] == num {
			ch <- true
			return
		}
	}
	return
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	c := make(chan bool)
	go search(arr[len(arr)/2:], 315, c)
	go search(arr[:len(arr)/2], 315, c)
	x, y := <-c, <-c
	fmt.Println(x || y)
	return
}
