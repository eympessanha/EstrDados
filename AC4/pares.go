/*Implemente em Go um algoritmo que resolva o seguinte problema: dado um array de números inteiros positivos,
considerado ordenado crescentemente, e um valor alvo, encontre um par de números no array cuja soma seja igual
 ao valor alvo. Se nenhum par for encontrado, retorne um valor (-1, -1) indicando que nenhum par foi encontrado.
  Resolva esse problema com um algoritmo cuja complexidade é O(n).*/

package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	alvo := 15

	fmt.Println(encontraPar(lista, alvo))
}

func encontraPar(lista []int, alvo int) (int, int) {
	for i := 0; i <= len(lista)-1; i++ {
		for j := 1; j <= len(lista); j++ {
			if i+j == alvo {
				return i, j
			}
		}
	}
	return -1, -1
}

// corrigir complexidade
