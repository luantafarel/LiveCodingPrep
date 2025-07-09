package main

import (
	"fmt"
	"time"
)

// Serviço de envio
func sender(id int, messages chan<- string) {
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Sender %d: Message %d", id, i)
		messages <- msg
		time.Sleep(time.Millisecond * 500) // Simula atraso no envio
	}
}

// Serviço de processamento
func processor(messages <-chan string, done chan<- bool) {
	for msg := range messages {
		fmt.Println("Processing:", msg)
		time.Sleep(time.Millisecond * 200) // Simula tempo de processamento
	}
	done <- true
}

func main() {
	messages := make(chan string)
	done := make(chan bool)

	// Inicia serviços de envio
	for i := 1; i <= 3; i++ {
		go sender(i, messages)
	}

	// Inicia o processador
	go processor(messages, done)

	// Espera o processamento terminar
	time.Sleep(time.Second * 3)
	close(messages)
	<-done
}
