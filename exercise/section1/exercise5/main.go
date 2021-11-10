package main

import "fmt"

/*
	Se deben ir agregando claves y valores por pantalla.
	Las claves y valores se deben agregar a un map:
	map[clave] = valor

	La clave debe ser de tipo numerico y el valor de tipo string

	Si ingreso una key que no es de tipo numerico, debe terminar la ejecucion del programa
	e imprimir por consola el map
*/

func main() {

	maps := make(map[int]string)
	fmt.Println("Seleccione valores, se sale con cero")

	for {
		var key int
		var value string

		fmt.Print("Seleccione key: ")
		if _, err := fmt.Scanf("%d", &key); err != nil {
			break
		}

		fmt.Print("Seleccione value: ")
		fmt.Scanf("%s", &value)

		maps[key] = value
	}

	fmt.Println("Los valores del map son: ", maps)
}
