package main

import (
	"fmt"
)

func main() {
	fmt.Println(uno(2))

	numero, status, _, _, _ := dos(1)
	fmt.Println(numero)
	fmt.Println(status)

	fmt.Println(tres(2))

	fmt.Println(calculo(1, 2, 3))

}

func uno(numero int) int {
	return numero * numero
}

func dos(numero int) (int, bool, int, int, int) {
	return numero, true, numero, numero, numero
}

func tres(numero int) (z int) {
	z = 20
	return
}

func calculo(numero ...int) int {
	total := 0
	for _, num := range numero {
		total += num
	}
	return total
}
