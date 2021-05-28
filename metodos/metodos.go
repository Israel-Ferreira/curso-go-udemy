package main

import "fmt"

type Usuario struct {
	nome string
	email string
	idade uint
}


func (u Usuario) ToString() string{
	return fmt.Sprintf("Nome: %s, Email: %s, Idade: %d", u.nome, u.email,u.idade)
}

func (u Usuario) Salvar(){
	fmt.Printf("Salvando os dados do Usuário %s no BD \n", u.nome)
}

func (u Usuario) IsMaiorDeIdade() bool{
	return u.idade >= 18
} 


func (u *Usuario) FazerAniversario() {
	u.idade++
}


func main() {
	user := Usuario{"Luke", "luke@jedi.com", 16}
	fmt.Println(user.ToString())

	fmt.Println(user.IsMaiorDeIdade())

	user.Salvar()


	user2 := Usuario{"Darth Vader", "darthvader@empire.com",55}
	fmt.Println(user2.ToString())

	user2.FazerAniversario()

	fmt.Println(user2.ToString())

	

	
}