package contato

import (
	"fmt"
)

/* Elabore um struct Contato, que deve conter os campos nome e e-mail. O struct deve
conter um método para alterar o e-mail.*/

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

/*Elabore uma função adicionarContato, que recebe um contato e um array com até 5 elementos e inclui o
contato no primeiro elemento vazio do array.*/

func AdicionarContato(contato Contato, listaContatos *[5]Contato) {
	for i := 0; i < len(listaContatos); i++ {
		if listaContatos[i].Nome == "" && listaContatos[i].Email == "" {
			listaContatos[i] = contato
			fmt.Println("Contato adicionado com sucesso!")
			return
		}
	}
	fmt.Println("Não foi possível adicionar o contato.")
}

/*Elabore uma função excluirContato, que recebe um array de 5 elementos e elimina
o último contato válido.*/

func ExcluirContato(listaContatos *[5]Contato) {
	for i := len(listaContatos) - 1; i >= 0; i-- {
		if listaContatos[i].Nome != "" || listaContatos[i].Email != "" {
			listaContatos[i] = Contato{} // limpar contato
			fmt.Println("Último contato excluído!")
			return
		}
	}
	fmt.Println("Não há contatos para excluir")
}
