package main

import "fmt"

func calculosMatemáticos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatemáticos(569, 856)
	fmt.Println(soma, subtracao)
}
