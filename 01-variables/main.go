package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	// declaramos una variable con var
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	// variable uint, solo numeros positivos
	var myUintVar uint
	myUintVar = 12
	fmt.Println("My variable is: ", myUintVar)

	// variable string
	var myStringVar string
	myStringVar = "My string variable"
	fmt.Println("My variable is: ", myStringVar)

	// variable booleana
	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable is: ", myBoolVar)

	// vemos la direccion de memoria de nuestra variable
	fmt.Println("My variable address is: ", &myStringVar)

	// go tomara el valor que asignamos como tipo de dato (al usar :=)
	myIntVar2 := 12
	fmt.Println("My variable is: ", myIntVar2)

	myStringVar2 := "My string variable with :="
	fmt.Println(myStringVar2)

	myBoolVar2 := true
	fmt.Println("My variable is: ", myBoolVar2)

	fmt.Println()

	// la constante tomara el valor asignado como tipo de dato
	const myFirstConst = "a12"
	fmt.Println("My const is: ", myFirstConst)

	// especificamos el tipo de dato de la constante
	const myIntConst int = 12
	fmt.Println("My const is: ", myIntConst)

	const myStringConst string = "aaa"
	fmt.Println("My const is: ", myStringConst)

	/*
		int8	8 bits	-128 to 127
		int16	16 bits	-2^15 to 2^15 -1
		int32	32 bits	-2^31 to 2^31 -1
		int64	64 bits	-2^63 to 2^63 -1
		int	Platform dependent	Platform dependent
	*/

	fmt.Println()

	var myVar int
	myVar = 23
	fmt.Printf("%T", myVar) // imprime tipo de dato (int)
	fmt.Printf("%d", myVar) // imprime valor (23)
	// %d int, %f float, %s string

	var my8BitsIntVar int8
	// valor por defecto 0
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)

	// unsafe.Sizeof nos devuelve el tamaño en bytes
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my32BitsIntVar int32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)

	var my64BitsIntVar int64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

	/*
		Type	Size	Range
		uint8	8 bits	0 to 255
		uint16	16 bits	0 to 2^16 -1
		uint32	32 bits	0 to 2^32 -1
		uint64	64 bits	0 to 2^64 -1
		uint	Platform dependent	Platform dependent
	*/

	fmt.Println()

	var my8BitsUnitVar uint8
	fmt.Printf("Uint default value: %d\n", my8BitsUnitVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsUnitVar, unsafe.Sizeof(my8BitsUnitVar), unsafe.Sizeof(my8BitsUnitVar)*8)

	var my16BitsUnitVar uint16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitsUnitVar, unsafe.Sizeof(my16BitsUnitVar), unsafe.Sizeof(my16BitsUnitVar)*8)

	var my32BitsUnitVar uint32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsUnitVar, unsafe.Sizeof(my32BitsUnitVar), unsafe.Sizeof(my32BitsUnitVar)*8)

	var my64BitsUnitVar uint64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsUnitVar, unsafe.Sizeof(my64BitsUnitVar), unsafe.Sizeof(my64BitsUnitVar)*8)

	fmt.Println()

	var myFloast32Var float32
	fmt.Printf("Float default value: %f\n", myFloast32Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloast32Var, unsafe.Sizeof(myFloast32Var), unsafe.Sizeof(myFloast32Var)*8)
	fmt.Println()

	var myFloat64Var float64
	fmt.Printf("Float default value: %f\n", myFloat64Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)
	fmt.Println()

	// string es una secuencia de bytes
	var myStringVar3 string
	fmt.Printf("String default value: %s\n", myStringVar3)

	myStringVar5 := `My string variable in golang
	with multiple
	line`

	fmt.Printf("The variable value is: %s\n", myStringVar5)

	// con las llaves definimos un scope
	{
		// conversiones
		fmt.Println()
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		intVal1, err := strconv.ParseInt("1333", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)

		intVal2, err := strconv.ParseInt("aa122", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal2, intVal2)

		floatVar1, _ := strconv.ParseFloat("-11.2", 64)
		fmt.Printf("type: %T, value: %f\n", floatVar1, floatVar1)

	}

	fmt.Println()
}
