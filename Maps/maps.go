package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Aline",
		"sobrenome": "Barbosa",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]string{
		"nome": {
			"primeiro": "Cristine",
			"ultimo": "Barbosa",
		}		
	}
	fmt.Println(usuario2)
}

