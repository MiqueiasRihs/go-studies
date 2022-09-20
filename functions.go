package main

import (
	"fmt"
)

var number = 10

func main() {

	second_number := 20

	showNumber(number)
	showNumber(second_number)
}

func showNumber(x int) {
	fmt.Println(x)
}
