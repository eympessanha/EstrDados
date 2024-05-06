package main

import "fmt"

func main() {
    var N int
    fmt.Scanln(&N)

    for i := 0; i < N; i++ {
        var comidaDisponivel float64
        fmt.Scanln(&comidaDisponivel)
        
        fmt.Println(dias(comidaDisponivel), "dias")
    }
    
}

func dias(comidaDisponivel float64) int {
    dias := 0
    for comidaDisponivel > 1 {
        dias += 1
        comidaDisponivel /= 2
    }
        
    return dias
}
