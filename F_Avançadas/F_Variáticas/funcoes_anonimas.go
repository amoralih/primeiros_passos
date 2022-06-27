package main

import "fmt"

//Exemplo 1 - Função Anônima

func main() {
	func(texto string) {
		fmt.Println("Olá Mundo!")
	}("Passando algum parâmetro")

}

// Exemplo 2 -  Função anônima com retorno

func receber() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando algum parâmetro")

	fmt.Println(retorno)

}
