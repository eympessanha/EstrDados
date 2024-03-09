package main

import 	"fmt"


func procedimentoHanoi(n int, origem, trabalho, destino string) {

	procedimentoHanoi(n-1, origem, destino, trabalho)

	// Mover o disco da origem para o destino
	procedimentoHanoi(n-1, trabalho, origem, destino)
}

func main() {
	n := 3
	origem, trabalho, destino := "A", "B", "C"

	fmt.Printf("Resolvendo a Torre de Hanói para %d discos\n", n)
	procedimentoHanoi(n, origem, destino, trabalho)
}