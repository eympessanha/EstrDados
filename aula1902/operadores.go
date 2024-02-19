package main

import "fmt"

func main(){
	var x = 2
	y := 4

	//operadores aritmeticos
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	//operadores de comparação
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)

	var a,b = true, false
	//operadores logicos
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!b)

	//operadores unarios
	x++
	y--
}