package main

import "fmt"

func main() {
	//aritmético
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var num1 int16 = 10
	var num2 int16 = 25
	soma2 := num1 + num2
	fmt.Println(soma2)

	//FIM DOS ARITMÉTICOS

	//ATRIBUIÇÃO

	var var1 string = "String"
	var2 := "String2"
	fmt.Println(var1, var2)
	//FIM ATRIBUIÇÃO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	//FIM RELACIONAIS

	//LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	//FIM LÓGICOS

	//OPERADORES UNÁRIOS
	num := 25
	num++
	num += 10
	fmt.Println(num)
	//FIM UNÁRIOS
}
