package main

import "fmt"

func evenSum(n []int) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		if n[i]%2 == 0 {
			sum += n[i]
		}
	}
	return sum
}

func main() {
	fmt.Println(evenSum([]int{1, 2, 3, 4, 5, 6, 10}))
	return
}
