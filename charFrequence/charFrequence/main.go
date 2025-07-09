package main

import (
	"fmt"
	"strings"
)

func charFrequency(s string) map[string]int {
	m := make(map[string]int)
	arrS := strings.Split(s, "")
	for _, val := range arrS {
		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}
	return m
}

func main() {
	fmt.Println(charFrequency("golang"))
	return
}
