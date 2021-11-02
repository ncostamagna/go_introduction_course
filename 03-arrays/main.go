package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar [6]int
	fmt.Println(myArrayVar)

	myArrayVar1 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar1)

	myArrayVar[1] = 2
	myArrayVar[2] = 5
	myArrayVar[3] = 9
	fmt.Println(myArrayVar)

	fmt.Println("Size: ", len(myArrayVar))
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)

	var slice1 []int
	fmt.Printf("size: %d, value %v\n", len(slice1), slice1)

	slice1 = append(slice1, 10, 20, 30, 40, 50)
	fmt.Printf("size: %d, value %v\n", len(slice1), slice1)

	mySlice := myArrayVar[2:4]
	fmt.Println(mySlice)
	fmt.Println("Size: ", len(mySlice))

	fmt.Println(&myArrayVar[2])
	fmt.Println(&mySlice[0])

	fmt.Println(myArrayVar)

	mySlice[0] = 80
	mySlice[1] = 100

	fmt.Println(myArrayVar)

	fmt.Println(myArrayVar[:4])
	fmt.Println(myArrayVar[2:])

	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Printf("size: %d, value %v\n", len(slice), slice)
}
