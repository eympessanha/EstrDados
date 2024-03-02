package main

import (
	"fmt"
	g "geometria/geometria"
	"math"
)

func main() {

	/*Faça um programa que cria um vetor de inteiros com 10 elementos. Então inicialize este vetor com números quaisquer
	  e imprima na tela todos os seus elementos, um abaixo do outro.*/
	VetorInteiros := [10]int{4, 8, 1, -5, 3, 9, 34, 93, 2, -7}

	fmt.Println("Valores da lista: ")
	for _, V := range VetorInteiros {
		fmt.Println(V)
	}


	var caracteres string
	fmt.Println("\nInforme uma palavra: ")
	fmt.Scanln(&caracteres)
	fmt.Println(InverteString(caracteres))


	ponto := Ponto{2, 9}
	fmt.Println("\nDistância do ponto 2,9: ", ponto.DistanciaOrigem())


	var base, altura float64
	fmt.Println("\nInforme o valor da altura do retângulo: ")
	fmt.Scanln(&altura)
	fmt.Println("\nInforme o valor da base do retângulo: ")
	fmt.Scanln(&base)

	ret := g.Retangulo{Base: base, Altura: altura}

	fmt.Println("\nA área do retangulo é ", ret.Area())
	fmt.Println("O perímetro do retangulo é ", ret.Perimetro())
}

/*Faça uma função/método que receba uma string como parâmetro e que
retorne uma nova string, onde a sequência dos caracteres foi invertida.
Dentro da parte principal (main), leia uma string digitada pelo usuário
e passe para a função/método criada, imprimindo em seguida a string devolvida.
 Para esse exercício, pesquise sobre o comportamento e a interação entre
 dados do tipo rune e do tipo string.*/

func InverteString(caracteres string) string {
	invertido := []rune(caracteres)
	for i, j := 0, len(invertido)-1; i < j; i, j = i+1, j-1 {
		invertido[i], invertido[j] = invertido[j], invertido[i]
	}
	return string(invertido)
}

/*Crie um tipo chamado Ponto que represente um ponto no plano cartesiano, com
coordenadas X e Y. Em seguida, implemente um método chamado DistanciaOrigem()
 que calcule a distância desse ponto até a origem (0,0).*/

type Ponto struct {
	X float64
	Y float64
}

func (p Ponto) DistanciaOrigem() float64 {
	return math.Sqrt((math.Pow(p.X, 2) + math.Pow(p.Y, 2)))
}
