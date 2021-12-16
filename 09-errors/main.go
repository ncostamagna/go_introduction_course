package main

import (
	"errors"
	"fmt"

	"github.com/ncostamagna/go_introduction_course/09-errors/operator"
)

func main() {

	var err error
	err = errors.New("my new error")
	fmt.Println(err)
	fmt.Println(err.Error())

	err2 := fmt.Errorf("my format err, string: %s, number: %d", "my string", 23)
	fmt.Println(err2)
	fmt.Println(err2.Error())

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("There aren't any results")
			fmt.Println("Recovered in ", r)
		}
	}()

	z := operator.Div(4, -2)

	fmt.Println("Results")
	fmt.Println("Z is", z)
}
