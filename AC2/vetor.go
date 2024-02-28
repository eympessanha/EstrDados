package main

import "fmt"

func main(){

/*Faça um programa que cria um vetor de inteiros com 10 elementos.
 Então inicialize este vetor com números quaisquer e imprima na tela todos os
 seus elementos, um abaixo do outro.*/
	vetorInteiros := [10]int{4, 8, 1, -5, 3, 9, 34, 93, 2, -7}

	fmt.Println("Valores da lista: ")
	for n := 0; n< len(vetorInteiros); n++{
		fmt.Println(vetorInteiros[n])
	}

	var palavra string
	fmt.Println("Informe uma palavra: ")
	fmt.Scanln(&palavra)
	fmt.Println(inverteString(palavra))



}

/*Faça uma função/método que receba uma string como parâmetro e que
 retorne uma nova string, onde a sequência dos caracteres foi invertida.
 Dentro da parte principal (main), leia uma string digitada pelo usuário
 e passe para a função/método criada, imprimindo em seguida a string devolvida.
  Para esse exercício, pesquise sobre o comportamento e a interação entre
  dados do tipo rune e do tipo string.*/

func inverteString(palavra string) string{
		invertido := []rune(palavra)
		for i, j := 0, len(invertido)-1; i < j; i, j = i+1, j-1 {
			invertido[i], invertido[j] = invertido[j], invertido[i]
		}
		return string(invertido)
}

//FALTA 3 E 4