package main

import (
	"fmt"
	"time"
)

//Primeiro tipo de FOR
func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)

	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	//Clausula Range

	nomes := [3]string{"Aline", "Cristine", "Barbosa"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	usuario := map[string]string{
		"nome":      "Fabio",
		"sobrenome": "Carito",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
