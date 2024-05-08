package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var M, P, quantidade int
		var produto string
		produtos := make(map[string]float64)

		fmt.Scan(&M)
		for j := 0; j < M; j++ {
			fmt.Scan(&produto)
			var preco float64
			fmt.Scan(&preco)
			produtos[produto] = preco
		}

		var total float64
		fmt.Scan(&P)
		for k := 0; k < P; k++ {
			fmt.Scan(&produto, &quantidade)
			total += float64(quantidade) * produtos[produto]
		}

		fmt.Printf("R$ %.2f\n", total)
	}
}

