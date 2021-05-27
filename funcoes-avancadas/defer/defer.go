package main

import "fmt"


func funcao1(){ 
	fmt.Println("Executando a função 1")
}

func funcao2(){
	fmt.Println("Executando a função 2")
}


func AlunoEstaAprovado(nota1, nota2 float64) (bool){
	defer fmt.Println("Média Calculada, O resultado será retornado")
	media := (nota1 + nota2) / 2

	if media >= 6 {
		return true
	}


	return false;

}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(AlunoEstaAprovado(10.0, 6.0))
}