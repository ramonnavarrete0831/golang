package main

import "fmt"

func main() {
	paises := make(map[string]string, 5)
	paises["mexico"] = "D.F"
	paises["argentina"] = "Buenos Aires"

	fmt.Println(paises)

	campeonato := map[string]int{
		"barcelona":   5,
		"real_madrid": 38,
		"chivas":      67,
	}

	campeonato["ramon"] = 2

	delete(campeonato, "real_madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Println(equipo, puntaje)
	}

	puntaje, existe := campeonato["chivas"]

	fmt.Println(puntaje, existe)

}
