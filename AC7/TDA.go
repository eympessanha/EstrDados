package main

import "fmt"

func main() {

	var casosDeTeste int
	fmt.Scanln(&casosDeTeste)

	for i := 1; i <= casosDeTeste; i++ {
		var n1, n2, d1, d2, numResultado, denResultado int
		var operador, barra1, barra2 string

		fmt.Scanln(&n1, &barra1, &d1, &operador, &n2, &barra2, &d2)

		if operador == "+" {
			numResultado, denResultado = somar(n1, d1, n2, d2)
		} else if operador == "-" {
			numResultado, denResultado = subtrair(n1, d1, n2, d2)
		} else if operador == "*" {
			numResultado, denResultado = multiplicar(n1, d1, n2, d2)
		} else {
			numResultado, denResultado = dividir(n1, d1, n2, d2)
		}

		numSimplificado, denSimplificado := simplificarFracao(numResultado, denResultado)

		fmt.Printf("%d/%d = %d/%d\n", numResultado, denResultado, numSimplificado, denSimplificado)

	}

}

func maximoDivisorComum(numerador, denominador int) int {
	for denominador != 0 {
		numerador, denominador = denominador, numerador%denominador
	}
	return numerador
}

func simplificarFracao(numerador, denominador int) (int, int) {
	divisorComum := maximoDivisorComum(numerador, denominador)
	// Aplica o sinal negativo apenas ao numerador, se necessário
	if numerador < 0 {
		return (numerador / divisorComum) * -1, (denominador / divisorComum) * -1
	}
	return numerador / divisorComum, denominador / divisorComum
}

func somar(n1, d1, n2, d2 int) (int, int) {
	numerador := n1*d2 + n2*d1
	denominador := d1 * d2
	return numerador, denominador
}

func subtrair(n1, d1, n2, d2 int) (int, int) {
	numerador := n1*d2 - n2*d1
	denominador := d1 * d2
	return numerador, denominador
}

func multiplicar(n1, d1, n2, d2 int) (int, int) {
	numerador := n1 * n2
	denominador := d1 * d2
	return numerador, denominador
}

func dividir(n1, d1, n2, d2 int) (int, int) {
	numerador := n1 * d2
	denominador := n2 * d1
	return numerador, denominador
}