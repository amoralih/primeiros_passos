package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivos structs")

	var user usuario
	user.nome = "Amora"
	user.idade = 34
	fmt.Println(user)

	enderecoExemplo := endereco{"Green Hill Zone", 6}

	user2 := usuario{"Amora", 34, enderecoExemplo}
	fmt.Println(user2)

	user3 := usuario{idade: 34}
	fmt.Println(user3)

}
