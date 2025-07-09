package main

import "fmt"

func fibonacciIterative(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	sum := 1
	prev := 1
	for i := 1; i < num; i++ {
		cup := sum
		sum += prev
		prev = cup
	}
	return sum
}

func fibonacciRecursive(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return fibonacciRecursive(num-1) + fibonacciRecursive(num-2)
}

func main() {
	fmt.Println(fibonacciIterative(3))
	fmt.Println(fibonacciRecursive(3))
	return
}

// 1 1 2 3 5 8

/* recursivo de 5 = recursivo de 3 + recursivo de 4
recursivo de 4 = recursivo de 2 + recursivo de 3
recursivo de 3 = recursivo de 1 + recursivo de 2
recursivo de 2 = recursivo de 0 + recursivo de 1
recursivo de 0 e 1 = 1
*/
