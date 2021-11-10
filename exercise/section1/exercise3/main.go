package main

import "fmt"

/*
	Se debe ir ingresando valores por consola hasta que se agrega
	cero y finaliza el programa, todos los valores deben ser numericos.

	En caso de agregar un valor que no sea numerico el programa lo ignorara y seguira
	pidiendo valores normalmente

	Todos los valores ingresados por consola se deben agregar a un array de tipo
	numerico e imprimirlo por pantalla al finalizar
*/

func main() {

	var array []int
	fmt.Println("Seleccione valores, se sale con cero")

	for {
		var value int

		if _, err := fmt.Scanf("%d", &value); err != nil {
			continue
		}

		if value == 0 {
			break
		}

		array = append(array, value)
	}

	fmt.Println("Los valores del array son: ", array)
}
