package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Fila struct {
	tamanho int
	inicio  *No
	fim    *No
}

func (f *Fila) percorre() {
	if f.inicio == nil {
		fmt.Println("Fila vazia")
	} else {
		no := f.inicio
		for no != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println()
	}
}

func (f *Fila) insere(novoValor int) {
	novoNo := &No{valor: novoValor}

	if f.fim == nil {
		f.inicio = novoNo
		f.fim = novoNo
	} else {
		f.fim.prox = novoNo
		f.fim = novoNo
	}

	f.tamanho++
}

func (f *Fila) remove() error {
	if f.inicio == nil {
		return fmt.Errorf("Fila vazia")
	}

	noRemovido := f.inicio
	f.inicio := f.inicio.prox

	if f.inicio == nil {
		f.fim = nil
	}

	f.tamanho--
	return noRemovido
}

func main() {
	fila := Fila{}

	fila.insere(2)
	fila.insere(4)
	fila.insere(8)

	fila.percorre()
	fmt.Println("Tamanho da fila:", fila.tamanho)

	fila.remove()

	fila.percorre()
	fmt.Println("Tamanho da fila:", fila.tamanho)

	fila.remove()
	fila.remove()

	fila.percorre()
	fmt.Println("Tamanho da fila:", fila.tamanho)

	err := fila.remove()
	fmt.Println(err)
}
