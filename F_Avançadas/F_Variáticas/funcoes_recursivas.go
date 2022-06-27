package main

import "fmt"

//Exemplo de Função Recursiva com Fibonacci
// 1 1 2 3 5 8 13 21...

func Fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return Fibonacci(posicao-2) + Fibonacci(posicao-1)
}

func main() {

	fmt.Println("Funções Recursivas")

	posicao := uint(10)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(Fibonacci(i))

	}

	fmt.Println(Fibonacci(posicao))

}
