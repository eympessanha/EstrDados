package main

import "fmt"

func main(){
	fmt.Println(calculaMedia(9,5,7,6))
	fmt.Println(verificaParidade(46))
	fmt.Println(verificaParidade(21))
	fmt.Println(minhaPotencia(2,3))
	fmt.Println(converteCelsiusParaFahrenheit(36))
}

/* Crie uma função calculaMedia que receba um número variável de parâmetros de tipo ponto flutuante e
retorne a média aritmética desses valores.*/

 func calculaMedia(valores ...float64) float64 {
	var soma float64
	var quantidade = float64(len(valores))
	for _, valor := range valores {
		soma += valor
	}
	var media = soma/quantidade

	return media
 }

 // Construa uma função verificaParidade que receba um inteiro e retorne se o número é par ou ímpar.

 func verificaParidade(numero int) string {
	if numero % 2 == 0 {
		return "Número par!"
	} else {
		return "Número ímpar!"
	}
 }

 /*Desenvolva uma função minhaPotencia que receba dois inteiros, base e expoente, e retorne o resultado de
 base elevado ao expoente, sem usar funções prontas da linguagem.*/

 func minhaPotencia(base, expoente int) int {
	resultado := 1
	if expoente == 0 {
		return 1
	} else {
		for i := 1; i <= expoente; i++{
			resultado *= base
		}
		return resultado
	}
 }

 /*Implemente uma função converteCelsiusParaFahrenheit que receba uma temperatura em Celsius e retorne a
 temperatura convertida em Fahrenheit.*/

 func converteCelsiusParaFahrenheit(celsius float64) float64{
	fahrenheit := (celsius * 9 / 5) + 32
	return fahrenheit
 }