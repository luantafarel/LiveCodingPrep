package main

import (
	"fmt"
	"slices"
	"strings"
)

// Exemplo: ["eat","tea","tan","ate","nat","bat"]
// Saída:   [["eat","tea","ate"], ["tan","nat"], ["bat"]]

// anagramChecker verifica se duas strings são anagramas
func anagramChecker(s, t string) bool {
	return sortedString(s) == sortedString(t)
}

// sortedString retorna a string com caracteres ordenados
func sortedString(s string) string {
	chars := strings.Split(s, "")
	slices.Sort(chars)
	return strings.Join(chars, "")
}

// anagramGrouper agrupa strings que são anagramas entre si
// Complexidade: O(n * k log k) onde n = número de strings, k = tamanho médio da string
func anagramGrouper(words []string) [][]string {
	// Mapa: chave ordenada -> lista de palavras originais
	groups := make(map[string][]string)

	for _, word := range words {
		key := sortedString(word)
		groups[key] = append(groups[key], word)
	}

	// Converte o mapa para slice de slices
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println("Input:", words)
	fmt.Println("Grouped anagrams:", anagramGrouper(words))

	// Teste individual
	fmt.Println("\n'eat' e 'tea' são anagramas?", anagramChecker("eat", "tea"))
	fmt.Println("'eat' e 'bat' são anagramas?", anagramChecker("eat", "bat"))
}

