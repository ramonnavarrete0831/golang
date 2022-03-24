package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var leyenda string
var resultado int

func main() {
	fmt.Println("Ingrese número 1:")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese número 2:")
	fmt.Scanln(&numero2)

	fmt.Println("Descripción:")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)

}
