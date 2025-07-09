package main

import "fmt"

type stackFunctions interface {
	Push() []int
	Pop() []int
}

type stack struct {
	items []int
}

func (t *stack) Push(n int) {
	t.items = append(t.items, n)
}

func (t *stack) Pop() {
	_, t.items = t.items[len(t.items)-1], t.items[:len(t.items)-1]
}

func main() {
	var s stack
	s.items = []int{1, 2, 3, 4}
	s.Push(6)
	fmt.Println(s.items)
	s.Pop()
	fmt.Println(s.items)
	return
}
