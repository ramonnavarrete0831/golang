package main

import "fmt"

var status bool

func main() {

	switch numero := 5; numero {
	case 1:
		fmt.Println(status)
	default:
		fmt.Println("Mayor que 1")
	}

	status = true
	if status = false; status == true {
		fmt.Println(status)
	} else {
		fmt.Println(status)
	}
}
