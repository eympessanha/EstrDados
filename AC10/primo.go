package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var X int
		fmt.Scanln(&X)

		if verificarPrimo(X) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}
	}
}


func verificarPrimo(X int) bool {
	if X <= 1 {
		return false
	} else if X <= 3 {
		return true
	} else if X % 2 == 0 || X % 3 == 0 {
		return false
	}
	
	for i := 5; i*i <= X; i += 6 {
		if X%i == 0 || X%(i+2) == 0 {
			return false
		}
	}
	return true
}
