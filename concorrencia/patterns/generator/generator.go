package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Marilene")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor Recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
