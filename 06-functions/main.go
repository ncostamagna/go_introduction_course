package main

import (
	"fmt"

	"github.com/ncostamagna/go_introduction_course/06-functions/function"
)

func main() {
	fmt.Println(function.Add(3, 4))

	function.RepeatString(10, "as")
	fmt.Println()
	v, err := function.Calc(function.SUM, 3, 6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value: ", v)
	}

	v, err = function.Calc(function.DIV, 3, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value: ", v)
	}

	x, y := function.Split(20)
	fmt.Println("Value1: ", x, " - Value2: ", y)

	fmt.Println()

	v = function.MSum(1, 2, 3)
	fmt.Println("multy sum: ", v)
	fmt.Println()

	v, err = function.MOperations(function.SUM, 2, 7, 1)
	fmt.Println("multy sum: ", v, " - error: ", err)

	v, err = function.MOperations(function.MUL, 2, 1, 3, 2, 1)
	fmt.Println("multy mul: ", v, " - error: ", err)

	v, err = function.MOperations(function.DIV, 2, 0, 1, 2, 1)
	fmt.Println("multy div: ", v, " - error: ", err)

	fmt.Println()

	fn := function.FactoryOperation(function.SUB)
	v = fn(2, 3)
	fmt.Println("Sub: ", v)

	fn = function.FactoryOperation(function.MUL)
	v = fn(2, 3)
	fmt.Println("Mul: ", v)
}
