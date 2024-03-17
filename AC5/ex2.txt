package main

import ( “fmt”
               “math” )

func main() {

var a, b, c float64

fmt.Println("Insira o valor de A, B e C: ")
fmt.Scanln(&a, &b, &c)

if a >= (b + c) {
	fmt.Println("NAO FORMA TRIANGULO")
} else if math.Pow(a,2) == (math.Pow(b,2) + math.Pow(c,2)) {
	fmt.Println("TRIANGULO RETANGULO")
} else if math.Pow(a,2) > (math.Pow(b,2) + math.Pow(c,2)) {
	fmt.Println("TRIANGULO OBTUSANGULO")
} else if math.Pow(a,2) < (math.Pow(b,2) + math.Pow(c,2)) {
	fmt.Println("TRIANGULO ACUTANGULO")
}

if (a == b) && (b == c) {
    fmt.Println("TRIANGULO EQUILATERO")
} else if (a == b) || (a == c) || (b == c) {
    fmt.Println("TRIANGULO ISOSCELES")}

}
