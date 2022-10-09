package main

import "fmt"

func main() {
	// Aritmeticos
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int8 = 10
	var numero2 int8 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Operadores lógicos
	fmt.Println("-------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Operadores Unários
	fmt.Println("-------------------")
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20
	fmt.Println(numero)

	numero *= 3 //numero = numero * 3
	fmt.Println(numero)

	fmt.Println("-------------------")
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
