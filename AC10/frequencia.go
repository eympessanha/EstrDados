package main

import (
        "fmt"
        "sort"
              )

func main() {
    var N int
    fmt.Scanln(&N)

    contador := make(map[int]int)
    
    for i := 0; i < N; i++ {
        var valor int
		fmt.Scanln(&valor)
		contador[valor]++
    }
    
    unicos := make([]int, 0, len(contador))
	for chave := range contador {
		unicos = append(unicos, chave)
	}
	sort.Ints(unicos)

	for _, valor := range unicos {
		fmt.Printf("%d aparece %d vez(es)\n", valor, contador[valor])
	}
}

