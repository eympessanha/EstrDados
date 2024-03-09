package AC4

func procedimentoHanoi(n int, origem, trabalho, destino string) {

	procedimentoHanoi(n-1, origem, destino, trabalho)

	// Mover o disco da origem para o destino
	procedimentoHanoi(n-1, trabalho, origem, destino)
}

