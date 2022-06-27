package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Aline",
		"sobrenome": "Barbosa",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":   "Cristine",
		"ultimo": "Barbosa",
		},
		"curso": {
			"nome": "Engenharia de Software",
			"campus": "Campus 2",
		},
		fmt.Println(usuario2)
		delete(usuario2, "nome")
		fmt.Println(usuario2)
	}
