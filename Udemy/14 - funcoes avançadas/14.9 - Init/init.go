package main

import "fmt"

// A funcao init é a funcao que vai ser executada antes de todas as outras
func init() {
	fmt.Println("Funçao init sendo executada")
}

func main() {
	fmt.Println("Funçao main sendo executada")
}
