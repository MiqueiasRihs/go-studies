package main

import "fmt"

func main() {
	func() {
		fmt.Println("Funcao anonima sem paremetro")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Funçao anonima com parametro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Retorno -> %s", texto)
	}("Funcao anonima com retorno")
	fmt.Println(retorno)
}
