package main

import  "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		var pa, pb int
		var g1, g2 float64
		fmt.Scanln(&pa, &pb, &g1, &g2)

		anos := 0

		for pa <= pb {
			pa = int(float64(pa) * (1 + g1/100))
			pb = int(float64(pb) * (1 + g2/100))
			anos++

			if anos > 100 {
				fmt.Println("Mais de 1 século")
				break
			}
		}

		if anos <= 100 {
			fmt.Printf("%d anos.\n", anos)
		}
	}
}
