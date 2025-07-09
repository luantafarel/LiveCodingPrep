package main

import "fmt"

func bubblesortAsc(arr []int) []int {
	for i := 0; i < len(arr)-1; {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
			i = 0
			continue
		}
		i++
	}
	return arr
}

func bubblesortDsc(arr []int) []int {
	for i := 0; i < len(arr)-1; {
		if arr[i] < arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
			i = 0
			continue
		}
		i++
	}
	return arr
}

func main() {
	arr := []int{1, 3, 3, 5, 14, 5, 6, 1, 2, 6, 5, 3}
	fmt.Println(bubblesortAsc(arr))
	fmt.Println(bubblesortDsc(arr))
	return
}
