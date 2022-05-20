package main

import "fmt"

func main() {
	var variable1 int = 20
	var variable2 int = variable1
	fmt.Println(variable1, variable2)

	variable1++
	fmt.Println(variable1, variable2)

	//O ponteiro é uma referência de memória
	var variable3 int
	var ponteiro *int

	variable3 = 100
	ponteiro = &variable3
	fmt.Println(variable3, ponteiro)

}
