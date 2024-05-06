package main

import ( 
        "fmt"
        "math"
              )

func main() {
    for {
        var X1,Y1,X2,Y2 int
        fmt.Scanln(&X1,&Y1,&X2,&Y2)
        
        if (X1 == 0 && Y1 == 0 && X2 == 0 && Y2 == 0) {
            break
        }
        
        fmt.Println(movimentosMinimos(X1,Y1,X2,Y2))
    }

}

func movimentosMinimos(X1, Y1, X2, Y2 int) int {
    if (X1 == X2 && Y1 == Y2) {
        return 0
    } else if (X1 == X2 || Y1 == Y2) {
        return 1
    } else if (math.Abs(float64(X1 - X2)) == math.Abs(float64(Y1 - Y2))) {
        return 1
    } else {
        return 2
    }
}
