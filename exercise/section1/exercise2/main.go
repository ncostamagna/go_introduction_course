package main

import "fmt"

/*
	Se debe ir ingresando valores por consola hasta que se agrega
	cero y finaliza el programa.

	Todos los valores ingresados por consola se deben agregar a un array
	e imprimirlo por pantalla al finalizar
*/

func main() {

	var array []string
	fmt.Println("Seleccione valores, se sale con cero")

	for {
		var value string
		fmt.Scanf("%s", &value)

		if value == "0" {
			break
		}

		array = append(array, value)
	}

	fmt.Println("Los valores del array son: ", array)
}
