package main

import "fmt"

func main() {
	// Ponteiro é uma referencia de memória
	var variavel1 int
	var ponteiro *int

	variavel1 = 100
	ponteiro = &variavel1

	fmt.Println(variavel1, ponteiro)
	fmt.Println(variavel1, *ponteiro) // Desreferenciação
}
