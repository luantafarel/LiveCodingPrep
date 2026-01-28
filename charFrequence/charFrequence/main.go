package main

import (
	"fmt"
	"strings"
)

func charFrequency(s string) map[string]int {
	m := make(map[string]int)
	arrS := strings.Split(s, "")
	for _, val := range arrS {
		m[val]++
	}
	return m
}

func main() {
	fmt.Println(charFrequency("golang"))
}
