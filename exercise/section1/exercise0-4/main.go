package main

import "fmt"

/*
	Dentro de nuestro codigo vamos a tener 2 variables definidas, que van a ser:
	- license: si la persona tiene licencia
	- age: la edad de la persona

	Si la persona tiene mas de 15 años y tiene licencia, debemos imprimir un mensaje que diga *Puede seguir avanzando*
	Si la persona tiene 15 años o menos, o no tiene licencia, debemos imprimir un mensaje que diga "No puede seguir circulando"

*/

func main() {

	license := true
	age := 15
	ok := false

	if age > 15 {
		if license == true {
			ok = true
		}
	}

	if ok {
		fmt.Println("Puede seguir avanzando")
	} else {
		fmt.Println("No puede seguir circulando")
	}
}
