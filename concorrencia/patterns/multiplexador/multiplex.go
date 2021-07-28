package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	multiplexador := multiplexar(escrever("Café"), escrever("Olá Mundo"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-multiplexador)
	}
}

func multiplexar(channelInput1, channelInput2 <-chan string) <-chan string {
	multiplexCh := make(chan string)

	go func ()  {
		for {
			select {
			case mensagem := <-channelInput1:
				multiplexCh <- mensagem
			case mensagem := <-channelInput2:
				multiplexCh <- mensagem
			}
		}
	}()



	return multiplexCh;
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor Recebido: %s", texto)
			time.Sleep(time.Duration(rand.Intn(20000)))
		}
	}()

	return canal
}
