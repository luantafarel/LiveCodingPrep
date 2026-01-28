package main

import "fmt"

// firstReverse inverte a string passada como parâmetro
func firstReverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Exemplos de teste da função firstReverse
	fmt.Println("Testando a função firstReverse:")
	fmt.Println()

	// Casos de teste
	testCases := []string{
		"hello",
		"world",
		"Go",
		"programming",
		"12345",
		"a",
		"",
		"racecar",
		"Hello World!",
		"áéíóú", // teste com acentos
	}

	for _, test := range testCases {
		result := firstReverse(test)
		fmt.Printf("firstReverse(\"%s\") = \"%s\"\n", test, result)
	}
}




