package main

func main() {
}

func buscaMatriz(m [][]int, n int, x int) bool {

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[i][j] == x {
				return true
			}
		}
	}
	return false
}

/*Dado um array de números inteiros positivos e um valor alvo, encontre um par de números no array
cuja soma seja igual ao valor alvo. Se nenhum par for encontrado, retorne um valor (-1, -1)
indicando que nenhum par foi encontrado*/
