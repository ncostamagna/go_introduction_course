package main

import "fmt"

const notebook string = "notebook"
const tv string = "tv"
const heladera string = "heladera"
const monitor string = "monitor"
const camara string = "camara"
const notFound string = "No encontrado"

/*
	Se deben ir agregando codigos por consola, los codigos son los siguientes:
	- "10" : "notebook"
	- "15" : "tv"
	- "21" : "heladera"
	- "27" : "monitor"
	- "35" : "camara"

	cada vez que voy agregando codigos por consola se debe ir agregando la descripcion a
	un array, si agrego un codigo que no existe se debe agregar el texto "No encontrado"
	al array

	Por ejemplo, si yo tipeo 10, 15 y 20 el array debe tener los valores
	["notebook", "tv", "No encontrado"]

	Si tipeo "0" se debe terminar la ejecucion del programa y debe mostrar por pantalla el
	array
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

		switch value {
		case "10":
			array = append(array, notebook)
		case "15":
			array = append(array, tv)
		case "21":
			array = append(array, heladera)
		case "27":
			array = append(array, monitor)
		case "35":
			array = append(array, camara)
		default:
			array = append(array, notFound)
		}

	}

	fmt.Println("Los valores del array son: ", array)
}
