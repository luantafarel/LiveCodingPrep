package main

import (
	"fmt"
	"slices"
	"strings"
)

// ["eat","tea","tan","ate","nat","bat"]

func anagramChecker(s, t string) bool {
	return sortedString(s) == sortedString(t)
}

func sortedString(s string) string {
	sortS := strings.Split(s, "")
	slices.Sort(sortS)
	return strings.Join(sortS, "")
}

func anagramGrouper(s []string) [][]string {
	m := make(map[string][]string)
	var r [][]string

	for i := 0; i < len(s); i++ {
		sI := sortedString(s[i])
		if _, ok := m[sI]; !ok {
			m[sI] = []string{s[i]}
		} else {
			m[sI] = append(m[sI], s[i])
		}
	}

	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println("Input:", words)
	fmt.Println("Grouped anagrams:", anagramGrouper(words))
}
