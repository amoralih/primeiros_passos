package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	pes1 := pessoa{"Amora", "Lih", 34, 160}
	fmt.Println(pes1)

	es1 := estudante{pes1, "Engenharia de Software", "Unicsul"}
	fmt.Println(es1)

}

//Em go você pode criar mais de uma struct e não usar, diferente das variáveis que o go reclama se você criar e não usar.
