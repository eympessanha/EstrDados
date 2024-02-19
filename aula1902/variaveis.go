/*
	inteiros
		int8, int16, int 32, int 64
		uint8, uint16, uint 32, uint 64 -> unsigned (sem sinal)
		int/uint -> generico, ocupa 32 ou 64 bits dependendo da arquitetura

		byte -> uint8
		rune -> int 32 -> caractere unicode

	ponto flutuante
		float32, float64(padrao)

		complex64 ou complex128 (i)

	texto
		string -> cada caracter ocupa 8 bits de memoria

	booleanos
		bool -> true, false

	outros tipos
		nil (nulos)
		ponteiros
		*/

package main

import "fmt"

const PI = 3.14
//z := 20 -> nao é ideal

func main(){
	var x int
	var msg1, msg2 string

	var y = 14.5 // ja interpreta o tipo
	z := 20 // nao é ideal

	msg := `texto
	em bloco`

	msg1 = "oi"
	msg2 = "ola"
	x = 34

	fmt.Println(msg)
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(x,y,z)
}