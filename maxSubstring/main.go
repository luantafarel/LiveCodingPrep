package main

import (
	"fmt"
)

func longestSubstringWithReplacement(s string, n int) (int, string) {
	charCount := make(map[byte]int)
	maxLen, maxCount, left := 0, 0, 0
	start, end := 0, 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]]++
		if charCount[s[right]] > maxCount {
			maxCount = charCount[s[right]]
		}

		// Verifica se a janela é válida
		if (right-left+1)-maxCount > n {
			charCount[s[left]]--
			left++
		}

		// Atualiza o tamanho máximo da substring e as posições de início e fim da janela
		if right-left+1 > maxLen {
			maxLen = right - left + 1
			start = left
			end = right
		}
	}

	// Reconstruindo a string completa com as mudanças
	mostFreqChar := s[start] // Assume o caractere mais frequente na janela como base
	for k, v := range charCount {
		if v > charCount[mostFreqChar] {
			mostFreqChar = k
		}
	}

	// Criando a string final com substituições
	finalString := ""
	for i := 0; i < len(s); i++ {
		if i >= start && i <= end && s[i] != mostFreqChar {
			finalString += string(mostFreqChar)
		} else {
			finalString += string(s[i])
		}
	}

	return maxLen, finalString
}

func main() {
	s := "AABABBA"
	n := 1
	maxLen, result := longestSubstringWithReplacement(s, n)
	fmt.Println("Tamanho da maior substring:", maxLen)
	fmt.Println("String completa com mudanças:", result)
}
