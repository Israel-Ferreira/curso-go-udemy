package main

import (
	"fmt"
	"funcoes/operacoes"
	"strings"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}


func calculosMatematicos(a, b int) (int, int, int, float64, error) {
	// soma := operacoes.
	soma := operacoes.Soma(a, b)
	sub := 	operacoes.Sub(a,b)
	multi := operacoes.Multi(a,b)

	div, err := operacoes.Div(a,b)

	if err != nil {
		return soma, sub, multi, 0, err
	}



	return soma,sub, multi, div, nil
}

func main() {
	resultadoSoma := somar(1,2)

	fmt.Println(resultadoSoma)

	f :=  func(txt string) string{
		fmt.Println(txt)

		return strings.ToUpper(txt)
	}



	textoMaiuscula := f("Qualquer Texto")
	fmt.Println(textoMaiuscula)

	a,_,c,d, error := calculosMatematicos(3,4)

	fmt.Println(a,c,d, error)
}