package main

import (
	"errors"
	"fmt"
)

/*
	Existe 4 tipos de numeros inteiros em Go, int8, int16, int32 e int64.
	A diferença entre eles é o tamanho que cada um suporta, onde int8
	armazena numeros menores e int64 armazena numeros muito maiores.
	Acontece semelhante tambem com float, onde temos o float32 e float64
*/

/*
	* uint - unsigned int *
	Unsigned significa sem sinal, não assinalado. Então é exatamente isso
	que ele é, um int que desconsidera o sinal.
*/

func main() {
	var numero1 int64 = 10000000000
	fmt.Println(numero1)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Float --------------------------
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	// End-Float --------------------------

	// Strings --------------------------
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)
	// End-Strings --------------------------

	// Valor inicial definido automaticamente quando nao atribuido nenhum outro valor.
	// No caso do 'texto' é 0, e no caso do 'texto2' é uma string vazia
	var texto int
	fmt.Println(texto)

	var texto2 string
	fmt.Println(texto2)

	// Bool --------------------------
	booleano2 := true
	fmt.Println(booleano2)

	var booleano1 bool
	fmt.Println(booleano1)
	// End-Bool --------------------------

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
