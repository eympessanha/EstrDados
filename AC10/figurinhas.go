package main

import "fmt"

func main() {
    var N int
    fmt.Scanln(&N)

    for i := 0; i < N; i++ {
        var n1, n2 int
        fmt.Scanln(&n1,&n2)
        
        resultado := tamanhoMax(n1, n2)
        fmt.Println(resultado)
    }

}

func mdc(n1, n2 int) int {
    for n2 != 0 {
        n1, n2 = n2, n1%n2
    }
    return n1
}

func tamanhoMax(n1, n2 int) int {
    divisorComum := mdc(n1, n2)
    return divisorComum
}
