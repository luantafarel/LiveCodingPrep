package main

import (
	"fmt"
	"slices"
	"strings"
)

func anagramChecker(s1, s2 string) bool {
	arrS1 := strings.Split(s1, "")
	arrS2 := strings.Split(s2, "")
	slices.Sort(arrS1)
	slices.Sort(arrS2)
	return strings.Join(arrS1, "") == strings.Join(arrS2, "")
}

func main() {
	fmt.Println(anagramChecker("jndbiosa", "dbiosajn"))
	return
}
