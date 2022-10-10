package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Encrementando i:", i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j += 5 {
	// 	fmt.Println("Incrementando j:", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Joao", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("--------------------------")

	for _, letra := range "MIQUÉIAS" {
		fmt.Println(string(letra))
	}

	fmt.Println("--------------------------")
	usuario := map[string]string{
		"nome":      "Miquéias",
		"Sobrenome": "Rihs",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, "=", valor)
	}
}
