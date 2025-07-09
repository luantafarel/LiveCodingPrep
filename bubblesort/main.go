package main

import "fmt"

func bubblesortAsc(num []int) []int {
	for i := 0; i < len(num)-1; {
		if num[i] > num[i+1] {
			temp := num[i]
			num[i] = num[i+1]
			num[i+1] = temp
			i = 0
			continue
		}
		i++
	}
	return num
}

func bubblesortDsc(num []int) []int {
	for i := 0; i < len(num)-1; {
		if num[i] < num[i+1] {
			temp := num[i]
			num[i] = num[i+1]
			num[i+1] = temp
			i = 0
			continue
		}
		i++
	}
	return num
}

func main() {
	fmt.Println(bubblesortAsc([]int{3, 4, 7, 1, 2, 9, 20, 15}))
	fmt.Println(bubblesortDsc([]int{3, 4, 7, 1, 2, 9, 20, 15}))
	return
}
