package main
 
import "fmt"
 
func main() {

    var C, P, F int
    fmt.Scanln(&C, &P, &F)

    totalPapel := C * F
    
    if P >= totalPapel {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
