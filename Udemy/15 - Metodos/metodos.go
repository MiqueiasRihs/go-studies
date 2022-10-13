package main

import "fmt"

// Metodo é obrigatoriamente associado á alguma estrura

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados.\n", u.nome)
}

func (u usuario) maiorDeIdade() {
	if u.idade < 18 {
		fmt.Println("Usuário menor de idade.")
	} else {
		fmt.Println("Usuário maior de idade.")
	}
}

/*
Quando voce tem algum metodo que irá fazer alguma alteracao no campo
do seu struct, o certo é passar a referencia do struct com (*)
*/
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Miquéias", 25}
	fmt.Println(usuario1)
	usuario1.salvar()
	usuario1.maiorDeIdade()

	usuario2 := usuario{"Ana", 15}
	usuario2.salvar()
	usuario2.maiorDeIdade()
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
