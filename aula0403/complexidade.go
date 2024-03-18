package main

func main() {
}

func buscaMatriz(m [][]int, n int, x int) bool {
	var i, j int
	i = 0

	for i < n {
		for j < n {
			if m[i][j] == x {
				return true
			}
			j++
		}
		i++
	}
	return false
}
