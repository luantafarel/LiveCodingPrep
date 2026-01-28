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
	m := make(map[string][]int)
	var r [][]string

	for i := 0; i < len(s); i++ {
		sI := sortedString(s[i])
		if _, ok := m[sI]; ok {
			// Já foi processado, pula
			continue
		}
		// Adiciona o índice atual
		m[sI] = append(m[sI], i)

		// Busca outros anagramas no resto do array
		for j := i + 1; j < len(s); j++ {
			if anagramChecker(s[i], s[j]) {
				m[sI] = append(m[sI], j)
			}
		}
	}

	// Converte os índices para as strings originais
	for _, v := range m {
		var group []string
		for _, i := range v {
			group = append(group, s[i])
		}
		r = append(r, group)
	}

	return r
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println("Input:", words)
	fmt.Println("Grouped anagrams:", anagramGrouper(words))
}

