package main
 
import "fmt"
 
func main() {

    var N int
    fmt.Scanln(&N)

    contador := 0
    i := 0
    V := 0
    
    for j := 0; j < N; j++ {
        fmt.Scanln(&V)
        
        if i != V {
            contador ++
        }
        
        i = V
    }
    
    fmt.Println(contador)

}
