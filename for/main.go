package main

import (
	"fmt"
)

func main() {
	fmt.Println("For")

	var i int = 0
RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a RUNINA sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("Valor de i : %d\n", i)
		i++
	}

	/*i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}*/

	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	/*numero := 0

	for {
		fmt.Println("Continuo")
		fmt.Scanln(&numero)

		if numero == 100 {
			break
		}
	}*/

	/*i := 1
	for i < 10 {
		fmt.Printf("\n Valor de i : %d", i)
		if i == 5 {
			fmt.Printf("	Multiplicamos por 2\n")
			i = i * 2
			continue
		}
		fmt.Printf("		Pasó por aquí\n")
		i++
	}*/

}
