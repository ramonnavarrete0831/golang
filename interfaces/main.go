package main

import "fmt"

type SerVivo interface{
	estaVivo() bool
}

/*Genero Humano*/

type humano interface{
	respirar()
	pensar()
	comer()
	sexo() string
}

type persona struct{
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool
	vivo bool
}

type hombre struct{
	persona
}

type mujer struct{
	persona
}

func(p *persona) respirar(){ p.respirando=true}
func(p *persona) comer(){ p.comiendo =true}
func(p *persona) pensar(){ p.pensando=true}
func(p *persona) sexo()string{ 
	if p.esHombre==true{
		return "Hombre"
	}else{
		return "Mujer"
	}
}
func(p *persona) estaVivo() bool { 
	return p.vivo
}



func HumanosRespidando(hu humano){
	hu.respirar();
	fmt.Printf("Soy un/a %s, y ya estoy respirando\n", hu.sexo())
}

/*Genero Animal*/

type animal interface{
	respirar()
	comer()
	EsCarnivoro() bool
}

type perro struct{
	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}

func(p *perro) respirar(){ p.respirando=true}
func(p *perro) comer(){ p.comiendo =true}
func(p *perro) EsCarnivoro() bool { return p.carnivoro}
func(p *perro) estaVivo() bool { 
	return p.vivo
}



func AnimalesRespirar(an animal){
	an.respirar()
	fmt.Println("Soy animal y estoy respirando")
}



/*type vegetal interface{
	ClasificacionVegetal() string
}*/



func AnimalesCarnovoros(an animal) int{
	if an.EsCarnivoro()==true{
		return 1
	}else{
		return 0
	}
}

/*SerVivo*/

func estoyVivo(sv SerVivo) bool{
	return sv.estaVivo()
}

func main() {
	fmt.Println("Interfaces Go")

	totalCarnivoros :=0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true

	AnimalesRespirar(Dogo)
	totalCarnivoros=+AnimalesCarnovoros(Dogo) 
	fmt.Printf("Total carnivoros %d\n",totalCarnivoros)
	fmt.Printf("Estoy vivo = %t\n",estoyVivo(Dogo))


	Pedro := new(hombre)
	Pedro.esHombre=true
	Pedro.vivo = false
	HumanosRespidando(Pedro)

	fmt.Printf("Estoy vivo = %t\n",estoyVivo(Pedro))

	Maria := new(mujer)
	Maria.vivo = true
	HumanosRespidando(Maria)

	fmt.Printf("Estoy vivo = %t\n",estoyVivo(Maria))
}
