package main
import "fmt"

func main() {
    for {
        var h1, m1, h2, m2, total int
        fmt.Scanln(&h1, &m1, &h2, &m2)
        
        if h1 == 0 && m1 == 0 && h2 == 0 && m2 ==0 {
            break
        }
        
        hora := (h1 * 60) + m1
        despertador := (h2 * 60) + m2
        
        if despertador > hora {
			total = despertador - hora
		} else {
			total = 1440 - hora + despertador
		}
        
        fmt.Println(total)
    }
}
