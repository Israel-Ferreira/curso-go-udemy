package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fibonacciNumbers :=  []uint{}
	posicao := uint(12)

	


	for i := uint(1); i <= posicao; i++ {
		fibonacciNumbers = append(fibonacciNumbers, fibonacci(i))
	}


	fmt.Println(fibonacciNumbers)

	
	fmt.Println(fibonacci(posicao))
}
