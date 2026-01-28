package main

import (
	"fmt"
)

func ReplaceInterval(sentence, replacement string, start, end int) string {
	return sentence[:start] + replacement + sentence[end:]
}

func main() {
	sentence1 := "Alice eats dinner with Bob at the movies."
	replacement1 := "Theater?"
	start1, end1 := 34, 41

	result1 := ReplaceInterval(sentence1, replacement1, start1, end1)
	fmt.Printf("Sentença: %q\n", sentence1)
	fmt.Printf("Palavra de substituição: %q\n", replacement1)
	fmt.Printf("Intervalo: [%d, %d]\n", start1, end1)
	fmt.Printf("Resultado: %q\n\n", result1)
}
