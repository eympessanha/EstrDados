package main

const M = 1000

type Pilha struct {
	pilha [M]rune
	topo  int
}

func (p *Pilha) exibirTopo() (rune, error) {

}

/*
	1069 beecrowd
	transformar entrada em pilha
	varrer a pilha para encontrar conjuntos
	quando encontrar, incrementar o contador e deletar os os caracteres da pilha (NÃO PODE???????)
*/
