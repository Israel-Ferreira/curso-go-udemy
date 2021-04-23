package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		Tipos Númericos em Go
		int8, int16, int32, int64 - Tipos de Númericos Inteiros; (byte, short, int, long)
		float32, float64  - Tipos Númericos de ponto flutuante (float, double)
	*/

	var numero int = -100
	fmt.Println(numero)


	var num2 uint = 200
	fmt.Println(num2)

	//alias
	// int32 = rune
	var num3 rune = 42
	fmt.Println(num3)

	// uint8 = byte
	var num4 byte = 123
	fmt.Println(num4)

	// Ponto Flutuante
	var nr1 float32 = 123.45
	var nr2 float64 = 123.80e6

	nr3 := 12345.60

	fmt.Printf("Nr1: %T, Nr2: %T, Nr3: %T \n", nr1, nr2, nr3)


	// Strings & Chars

	str1, str2 := "Café","Texto"

	fmt.Println(str1, str2)

	char := 'B'

	fmt.Printf("%d \n", char)


	var bool1 bool = true
	bool2 := false

	fmt.Println(bool1, bool2)


	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
