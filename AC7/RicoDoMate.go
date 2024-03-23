package main

import "fmt"

// n ta completo

 func main() {
	//quantidade de pessoas na roda
	var n int
	//litros de agua na garrafa termica e na cuia, respectivamente
	var l, q float64

	fmt.Scanln(&n, &l, &q)

	participantes := make([]string, n)

	// separar participantes
	for i := 0; i < n; i++ {
		fmt.Scan(&participantes[i])
	}

	//indice do ultimo participante a beber na roda
	indiceUltimo:= int(l/q) % n

	//quanto a ultima pessoa vai beber
	restante := l
	for restante != 0 {
		if restante - q < 0 {
			break
		}
		restante -= q
	}

	fmt.Printf("%s %.1f\n", participantes[indiceUltimo], restante)

}