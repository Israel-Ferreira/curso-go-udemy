package main

import "fmt"

func closure() func() {

	text := "Dentro da função closure"

	fnc := func() {
		fmt.Println(text)
	}

	return fnc
}

func main() {
	text := "Dentro da função main"
	fmt.Println(text)

	funcaoNova := closure()
	funcaoNova()
}
