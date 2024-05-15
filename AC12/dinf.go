package main
 
import (
        "fmt"
        "math"
              )
 
func main() {
    var n, h, c, l int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		fmt.Scan(&h, &c, &l)

		baseRampa := c * n
		alturaRampa := h * n
		larguraRampa := l

		comprimentoRampa := math.Sqrt(math.Pow(float64(alturaRampa), 2) + math.Pow(float64(baseRampa), 2))

		areaRampa := float64(larguraRampa) * comprimentoRampa

		fmt.Printf("%.4f\n", areaRampa/10000)
	}

}
