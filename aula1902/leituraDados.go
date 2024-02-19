package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var x, y, z int

	fmt.Println("Insira um valor")
	fmt.Scanln(&x)
	fmt.Println(x)

	fmt.Println("Insira dois valores")
	fmt.Scanln(&y, &z)
	fmt.Println(y, z)

	var nome string
	fmt.Println("Informe um nome")
	fmt.Scanln(&nome)
	fmt.Println(nome)

	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("Informe um nome")
	nome, _ = leitor.ReadString('\n')
	nome = strings.ReplaceAll(nome, "\n", "")
	fmt.Println(nome)
}
