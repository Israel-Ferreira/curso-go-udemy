package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Valor Zero em Arrays = []
	var array1 [5]int

	fmt.Println(array1)

	for i := 5; i > 0; i-- {
		array1[i-1] = i
	}

	fmt.Println(array1)

	array3 := [...]string{"Um", "Dois", "Três", "Quatro", "Cinco", "Seis"}

	fmt.Println(len(array3))

	sliceDeInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	/*
		Adicionando um valor no slice com a função append que retorna a copia do slice com o item informado pelo segundo argumento.
	*/
	sliceDeInt = append(sliceDeInt, 14)

	fmt.Println(sliceDeInt)
	fmt.Println(reflect.TypeOf(sliceDeInt))

	fmt.Println(reflect.TypeOf(array3))

	/* 
	Criando um slice a partir de um array
	*/

	slice1 := array1[:3]
	fmt.Println(slice1)

	array1[2] = 5

	fmt.Println(slice1)


	// Arrays Internos
	slice3 := make([]float64, 10, 11)
	fmt.Println(slice3)

	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	fmt.Println(slice3)

	slice3 = append(slice3, 12)
	slice3 =  append(slice3, 13)

	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	fmt.Println(slice3)


	slice4 := make([]float32, 5)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacidade

	fmt.Println(slice4)

	slice4 = append(slice4, 123)

	fmt.Println(slice4)

}
