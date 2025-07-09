package main

import "fmt"

// Etapa 1: Gerar números de 1 a 10
func generateNumbers(out chan<- int) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	close(out)
}

// Etapa 2: Dobrar os valores
func doubleNumbers(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

// Etapa 3: Filtrar múltiplos de três
func filterMultiplesOfThree(in <-chan int, out chan<- int) {
	for num := range in {
		if num%3 == 0 {
			out <- num
		}
	}
	close(out)
}

func main() {
	// Canais intermediários
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Iniciar as Goroutines do pipeline
	go generateNumbers(ch1)
	go doubleNumbers(ch1, ch2)
	go filterMultiplesOfThree(ch2, ch3)

	// Consumir o resultado final
	for result := range ch3 {
		fmt.Println(result)
	}
}
