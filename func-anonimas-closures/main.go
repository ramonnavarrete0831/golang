package main

import "fmt"

var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	tablaDel2 := 2
	Mitabla := Tabla(tablaDel2)
	for i := 1; i < 11; i++ {
		fmt.Println(Mitabla())
	}

	fmt.Printf("Sumo 5 + 7 = %d\n", calculo(5, 7))

	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Sumo 5 - 7 = %d\n", calculo(5, 7))

	calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}

	fmt.Printf("Multiplico 5 * 7 = %d\n", calculo(5, 7))

	operaciones()

}

func operaciones() {
	result := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(result())
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
