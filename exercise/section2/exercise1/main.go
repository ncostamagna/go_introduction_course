package main

import (
	"fmt"

	"github.com/ncostamagna/go_introduction_course/exercise/section2/exercise1/matrix"
)

/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que:

 - Reciba una serie de valores de punto flotante e inicialice los valores en la estructura Matrix
 - Tenga un metodo Print que imprime por pantalla la matriz de una formas más visible

La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho y si es cuadrática.
*/

func main() {
	m, err := matrix.New([]float64{2, 7, 8}, []float64{4, 4, 7}, []float64{5, 6, 1})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m.Print()
}
