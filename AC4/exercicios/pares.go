/*Implemente em Go um algoritmo que resolva o seguinte problema: dado um array de números inteiros positivos,
considerado ordenado crescentemente, e um valor alvo, encontre um par de números no array cuja soma seja igual
 ao valor alvo. Se nenhum par for encontrado, retorne um valor (-1, -1) indicando que nenhum par foi encontrado.
  Resolva esse problema com um algoritmo cuja complexidade é O(n).*/

package AC4

func encontraPar(lista []int, alvo int) (int, int) {
	i := 0
	j := 1
	for j <= len(lista) {
		soma := 


  for ponteiro_esquerdo < ponteiro_direito {
    soma := array[ponteiro_esquerdo] + array[ponteiro_direito]

    if soma == desejado {
      return ponteiro_esquerdo, ponteiro_direito
    } else if soma < desejado {
      ponteiro_esquerdo++
    } else {
      ponteiro_direito--
    }
  }

  return -1, -1
}

	
}

// corrigir complexidade
