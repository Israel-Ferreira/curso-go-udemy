package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup;

	waitGroup.Add(2)

	go func ()  {
		escrever("Hello World")
		waitGroup.Done()
	}()

	


	go func(){
		escrever("Programando em go!")
		waitGroup.Done()
	}()

	
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
