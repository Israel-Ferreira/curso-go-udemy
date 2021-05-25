package main

import "fmt"


func CalculosMatematicos(x, y int) (soma, sub int) {
	soma = x + y
	sub = x - y 

	return
}


func main() {
	soma, sub := CalculosMatematicos(13,40)
	fmt.Printf("Soma: %d, Substração: %d \n", soma, sub)
}
