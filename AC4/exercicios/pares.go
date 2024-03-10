/*Implemente em Go um algoritmo que resolva o seguinte problema: dado um array de números inteiros positivos,
considerado ordenado crescentemente, e um valor alvo, encontre um par de números no array cuja soma seja igual
 ao valor alvo. Se nenhum par for encontrado, retorne um valor (-1, -1) indicando que nenhum par foi encontrado.
  Resolva esse problema com um algoritmo cuja complexidade é O(n).*/

package AC4

func encontraPar(lista []int, alvo int) (int, int) {
	i := 0
	j := len(lista) - 1
	for i<j {
		soma := lista[i] + lista[j]

    if soma == alvo {
      return i,j
    } else if soma < alvo {
      i++
    } else {
      j--
    }
  }
  return -1, -1
}

	
