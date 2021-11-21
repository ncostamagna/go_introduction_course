package function

import (
	"errors"
	"fmt"
)

var myPrivateVar = 3

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(x, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Print(value)
	}
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y mustn't be zero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil

	}

	return 0, errors.New("Has been an error")
}

func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}

	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0, errors.New("There aren't values")
	}
	sum := values[0]

	for _, v := range values[1:] {

		switch op {
		case SUM:
			sum += v
		case SUB:
			sum -= v
		case DIV:
			if v == 0 {
				return 0, errors.New("mustn't be zero values")
			}
			sum /= v
		case MUL:
			sum *= v
		}
	}
	return sum, nil
}
