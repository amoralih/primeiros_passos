package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Uma função com dois retornos
func calculosMatematicos(a1, a2 int8) (int8, int8) {
	soma := a1 + a2
	subtracao := a1 - a2
	return soma, subtracao
}

func main() {
	soma := somar(6, 15)
	fmt.Println(soma)

	var milho = func(txt string) {
		fmt.Println(txt)
	}
	milho("Texto da função milho")

	resultadosoma, resultadosubtracao := calculosMatematicos(30, 67)
	fmt.Println(resultadosoma, resultadosubtracao)
}
