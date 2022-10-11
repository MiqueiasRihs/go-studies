package main

import "fmt"

func funcao1() {
	fmt.Println("Funcao 1")
}

func funcao2() {
	fmt.Println("Funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retonado!")
	fmt.Println("Funçao para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}

	return false
}

// Defer adia a execuçao de uma funcao até o ultimo momento possivel
func main() {
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoEstaAprovado(6, 6))

}
