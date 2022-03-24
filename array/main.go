package main

import "fmt"

var tabla [10]int
var matriz [5][7]int

func main() {

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	variante2()

	variante3()

	variante4()
	return

	matriz[0][5] = 2
	fmt.Println(matriz)

	tabla[5] = 5
	fmt.Println(tabla)

	table := [10]int{5, 52, 54, 2, 3, 1, 64, 39, 25, 62}
	fmt.Println(table)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}

	for _, element := range table {
		fmt.Println(element)
	}
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))

}
