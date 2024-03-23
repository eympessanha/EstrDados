package main

import "fmt"

func main(){

	var codigoPeca1, codigoPeca2, numeroPecas1, numeroPecas2 int
	var valorUnitarioPeca1, valorUnitarioPeca2 float64

	fmt.Scanln(&codigoPeca1, &numeroPecas1, &valorUnitarioPeca1)
	fmt.Scanln(&codigoPeca2, &numeroPecas2, &valorUnitarioPeca2)

	total := (float64(numeroPecas1) * valorUnitarioPeca1) + (float64(numeroPecas2) * valorUnitarioPeca2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)

}