package main

import (
	"fmt"
)

func main() {
	yearsOld := 32

	fmt.Println("Operators")
	fmt.Println(yearsOld > 30)  // true
	fmt.Println(yearsOld < 32)  // false
	fmt.Println(yearsOld <= 32) // true
	fmt.Println(yearsOld >= 40) // false
	fmt.Println(yearsOld == 32) // true

	fmt.Println()
	// OR

	fmt.Println(yearsOld < 32 || yearsOld == 32) // (false or true) = true
	fmt.Println(yearsOld < 32 || yearsOld == 33) // (false or false) = false
	fmt.Println(yearsOld < 40 || yearsOld == 33) // (true or true) = true

	fmt.Println()
	// AND

	fmt.Println(yearsOld < 32 && yearsOld == 32) // (false and true) = false
	fmt.Println(yearsOld < 32 && yearsOld == 33) // (false and false) = false
	fmt.Println(yearsOld < 40 && yearsOld == 32) // (true and true) = true

	fmt.Println()
	// NOT

	fmt.Println(true)             // true
	fmt.Println(!true)            // false
	fmt.Println(yearsOld < 40)    // true
	fmt.Println(!(yearsOld < 40)) // false

	// combinación de operaciones logicas
	fmt.Println(yearsOld < 25 && yearsOld == 32 || yearsOld < 40)   // (false and true or true ) = true
	fmt.Println(yearsOld < 25 && (yearsOld == 32 || yearsOld < 40)) // (false and (true or true) ) = false

	// condicional if
	yearsOld = 20

	if yearsOld > 18 {
		fmt.Printf("%d is higher than 18\n", yearsOld)
	}

	boolVal := false

	if boolVal {
		fmt.Println("is true")
	} else {
		fmt.Println("is false")
	}

	if value := true; value {
		fmt.Println("is true")
	}

	number := 3

	if number == 1 {
		fmt.Println("one")
	} else if number == 2 {
		fmt.Println("two")
	} else if number == 3 {
		fmt.Println("three")
	}

	// condicional switch
	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("undefined number")
	}

	switch number := 1; number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("undefined number")
	}

	switch {
	case number > 0:
		fmt.Println("positive")
	case number < 0:
		fmt.Println("negative")
	case number == 0:
		fmt.Println("zero")
	}
}
