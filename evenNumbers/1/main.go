package main

import "fmt"

func main() {
	chan := make (chan int, 20)
	go func () {
		for i:=0; i<200; i+=2 {
			chan <- i
		}
	}
	return
}
