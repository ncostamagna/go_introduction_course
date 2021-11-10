package main

import "fmt"

/*
	Dentro de nuestro codigo ya tenemos definido un array
	1) Debemos incrementar todos sus valores en 20
*/

func main() {

	array := [5]int{4, 2, 5, 6, 7}

	for i := range array {
		array[i] += 20
	}

	fmt.Println("Los valores del array son: ", array)
}
