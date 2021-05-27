package main

import "fmt"

func recuperarExecucao(){
	if r := recover(); r != nil {
		fmt.Println("Função Recuperada com sucesso")
	}
}

func AlunoEstaAprovado(n1, n2 float64) bool {
	media := (n1 + n2) / 2

	defer fmt.Printf("A media do aluno é: %.2f \n ", media)

	defer recuperarExecucao()

	if media >= 6 && media <= 10 {
		return true
	} else if media <= 6 && media >= 0 {
		return false
	}

	panic("A Media deve ser de (0 a 10)")
}

func main() {
	aprovado := AlunoEstaAprovado(-1.0, -4.0)
	fmt.Println(aprovado)
}
