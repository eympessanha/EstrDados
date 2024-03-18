package main

import "fmt"

func main() {
	tomadas()
	meteoritos()
	maiorNumero()
}

func tomadas() {
	var t1, t2, t3, t4 int8
	fmt.Scanln(&t1, &t2, &t3, &t4)
	fmt.Println(t1 + t2 + t3 + t4 - 3)
}

func meteoritos() {
	var x1, x2, x, y1, y2, y, n int
	num := make([]int, 0)
	i := 0
	for {
		fmt.Scanln(x1, x2, y1, y2)
		if x1 == 0 && x2 == 0 && y1 == 0 && y2 == 0 {
			break
		}
		fmt.Scanln(&n)
		num = append(num, 0)
		for j := 1; j <= n; j++ {
			fmt.Scanln(&x, &y)
			if x1 <= x && x <= x2 && y2 <= y && y <= y1 {
				num[i]++
			}
		}
		i++
	}
	for j, num_meteoritos := range num {
		fmt.Println("Teste", j+1)
		fmt.Println(num_meteoritos)
		fmt.Println("")
	}
}

func maiorNumero() {
	var num int
	maior := 0

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num > maior {
			maior = num
		}
	}
	fmt.Println(maior)
}
