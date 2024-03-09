package main

import (
	"fmt"
	e "AC4/exercicios"
)

func main() {

	// ex1
	n := 3
	origem, trabalho, destino := "A", "B", "C"

	fmt.Printf("Resolvendo a Torre de Hanói para %d discos\n", n)
	e.procedimentoHanoi(n, origem, destino, trabalho)

	//ex2
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	alvo := 15

	fmt.Println(e.encontraPar(lista, alvo))
}
