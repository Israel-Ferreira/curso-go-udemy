package main

import "fmt"

func main() {

	//funções anônimas
	newStr := func (text string) string {
		return fmt.Sprintf("Parametro Recebido -> %s", text)
	}("Israel")

	fmt.Println(newStr)
}