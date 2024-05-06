package main

import "fmt"

func main() {
    for {
        var M, N int
        
		_, err := fmt.Scanln(&M, &N)
		if err != nil {
			break 
		}
		
		fmt.Println(fatorial(M) + fatorial(N))
    }
}

func fatorial(M int) int {
    if M == 0 {
        return 1
    } else {
        fat := 1
        for i := 1; i <= M; i++ {
            fat *= i
        }
        return fat
    }
}
