package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return: "Domingo"
	case 2:
		return: "Segunda-Feira"
	case 3:
		return: "Terça-Feira"
	case 4:
		return: "Quarta-Feira"
	case 5:
		return: "Quinta-Feira"
	case 6:
		return: "Sexta-Feira"
	case 7:
		return: "Sábado"
	default:
		return: "Número Inválido"	 

	}
}

func diadiaDaSemana2(numero int) string {
	switch{
	case numero ==1
		return:"Domingo"
	case numero ==2
		return:"Segunda"
	case numero ==3
		return:"Terça"
	case numero ==4
		return:"Quarta"
	case numero ==5
		return:"Quinta"

	}
}

func main() {
	fmt.Println("Usando Switch")

	dia = didiaDaSemana(3)
	fmt.Println(dia)
}