package main

import (
	"fmt"
)

var y = "olá bom dia"

func main() {

	x := 10

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)
	// := Operador de declaracao, com ele é defenida pelo menos uma nova variavel
	// e ela é inserida apenas no code blocks
	// Ex: x := 20 nao irá funcionar pois a variavel x ja foi delacarada a cima
	// Então com ela podemos apenas fazer uma atribuição. Ex: x = 20

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T", z, z)

}
