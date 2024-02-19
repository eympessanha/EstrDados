package main

import "fmt"

func main(){
	// estruturas de decisão

	//if/else
	idade := 12
	if idade >= 18 {
		fmt.Println("Entrada liberada")
	} else if idade < 0 {
		fmt.Println("Erro de dados")
	} else {
		fmt.Println("Entrada proibida")
	}

	// switch case (não tem break)
	diaSemana := "segunda"

	switch diaSemana {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("Dia útil")
	case "sábado", "domingo":
		fmt.Println("Final de semana")
	default:
		fmt.Println("Dia inválido")
	}

	switch {
	case idade >= 18:
		fmt.Println("Entrada liberada")
	case idade < 0:
		fmt.Println("Dado inválido")
	default:
		fmt.Println("Entrada proibida")
	}

	// estruturas de decisao -> for
	//comportamento do while
	var x = 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	//comportamento do for
	for i := 0; i< 10; i++{
		if i == 3 { continue }
		if i == 4 { break }
		fmt.Println(i)
	}

}