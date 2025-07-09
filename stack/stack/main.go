package main

import "fmt"

type stackFunctions interface {
	Push()
	Pop()
	Print()
}

type stack struct {
	val []int
}

func (s *stack) Push(val int) {
	s.val = append(s.val, val)
}

func (s *stack) Pop() {
	s.val = s.val[:len(s.val)-1]
}

func (s *stack) Print() {
	for i := len(s.val) - 1; i >= 0; i-- {
		fmt.Printf("%d \n", s.val[i])
	}
	fmt.Println()
}

func main() {
	s := stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Print()
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println("After Popping 3 times")
	s.Print()
	return
}
