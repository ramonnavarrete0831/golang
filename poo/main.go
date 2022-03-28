package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.Usuario
}

func main() {
	user := new(pepe)
	user.AltaUsuario(1, "Pepe", time.Now(), true)
	fmt.Println(user)

	/*User := new(usuario)
	User.Nombre = "Ramón"
	fmt.Println(User)*/

}
