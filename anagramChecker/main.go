package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagramChecker(s1, s2 string) bool {
	s1Arr := strings.Split(s1, "")
	s2Arr := strings.Split(s2, "")
	sort.Strings(s1Arr)
	sort.Strings(s2Arr)
	return strings.Join(s1Arr, "") == strings.Join(s2Arr, "")
}

func main() {
	fmt.Println(anagramChecker("golang", "galong"))
	return
}
