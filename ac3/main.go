package main

/*Elabore uma interface por linha de comando na função main, que cria um array de 5 elementos e
permite a inclusão ou exclusão de contatos.*/

import (
	c "ac3/contato"
	"fmt"
)

func main() {
	contato1 := c.Contato{Nome: "Esther", Email: "esther@email.com"}
	contato2 := c.Contato{Nome: "Junior", Email: "junior@email.com"}

	var listaContatos [5]c.Contato

	c.AdicionarContato(contato1, &listaContatos)
	c.AdicionarContato(contato2, &listaContatos)

	fmt.Println("Lista de contatos: ")
	for _, contato := range listaContatos {
		fmt.Printf("Nome: %s, Email: %s\n", contato.Nome, contato.Email)
	}

	listaContatos[0].AlterarEmail("novoesther@email.com")

	c.ExcluirContato(&listaContatos)

	fmt.Println("\nLista de contatos após exclusão:")
	for _, contato := range listaContatos {
		fmt.Printf("Nome: %s, Email: %s\n", contato.Nome, contato.Email)
	}
}
