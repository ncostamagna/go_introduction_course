package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := strconv.Itoa(-42)
	fmt.Println(s)

	s = strconv.FormatBool(true)
	fmt.Println(s)
	s = strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s)
	s = strconv.FormatInt(-42, 10)
	fmt.Println(s)
	s = strconv.FormatUint(42, 10)
	fmt.Println(s)

	b, err := strconv.ParseBool("sarasa")
	fmt.Println(b, err)

	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f, err)

	in, err := strconv.ParseInt("-42t", 10, 64)
	fmt.Println(in, err)

	u, err := strconv.ParseUint("42", 10, 64)
	fmt.Println(u, err)
}
