package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	var u usuario
	u.nome = "Miquéias"
	u.idade = 25
	fmt.Println(u)

	exEndereco := endereco{"Rua Santa Maria", 0}

	usuario2 := usuario{"Miquéias", 25, exEndereco}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Rihs", endereco: exEndereco}
	fmt.Println(usuario3)
	fmt.Println(usuario3.endereco.logradouro)

}
