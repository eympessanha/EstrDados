package geometria

/*Crie um pacote chamado geometria que contenha funções
para calcular a área e o perímetro de um retângulo. Em seguida, use
esse pacote em um programa principal para calcular a área e o perímetro
de um retângulo com dimensões fornecidas pelo usuário.*/

type Retangulo struct{
	Base float64
	Altura float64
}

func (r Retangulo) Perimetro() float64 {
	return (r.Base*2) + (r.Altura*2)
}

func (r Retangulo) Area() float64 {
	return r.Base * r.Altura
}