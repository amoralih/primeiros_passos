package main

import (
	"errors"
	"fmt"
)

func main() {
	//números inteiros
	//A diferença é a quantidade de bits que cada carrega.
	//int8, int16, int32, int64
	var numero int16 = 100
	fmt.Println(numero)
	num := 10000000
	fmt.Println(num)

	//alias - apelidos para usar no Go
	//int32 = rune
	//uint8 = byte
	var numer rune = 12345
	fmt.Println(numer)
	var numeral byte = 123
	fmt.Println(numeral)

	//números reais
	//float32, float64

	var numreal float32 = 787.90
	fmt.Println(numreal)
	var numeroreal float64 = 1890000.89
	fmt.Println(numeroreal)

	//FIM NUMEROS

	//texto

	var str string = "sonic"
	fmt.Println(str)

	str2 := "tails"
	fmt.Println(str2)

	//FIM STRING

	//booleano
	var agora bool = true
	fmt.Println(agora)
	var passado bool = false
	fmt.Println(passado)

	//FIM BOOLEANO

	//erro

	var erro error
	fmt.Println(erro)

	var errado error = errors.New("Erro interno")
	fmt.Println(errado)

}
