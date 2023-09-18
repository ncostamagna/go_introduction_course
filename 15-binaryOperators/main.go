package main

import "fmt"

func main() {
	var a uint16 = 213

	fmt.Printf("Number %d in binary is %b\n", a, a)

	fmt.Printf("\nLeft shift\n")
	fmt.Printf("Number %d in binary is %b\n", a<<1, a<<1)
	fmt.Printf("Number %d in binary is %b\n", a<<10, a<<10)

	fmt.Printf("\nRight shift\n")
	fmt.Printf("Number %d in binary is %b\n", a>>1, a>>1)
	fmt.Printf("Number %d in binary is %b\n", a>>10, a>>10)
	fmt.Printf("Number %d in binary is %b\n", a>>5, a>>5)

	fmt.Printf("\n'a' and 'b' value:\n")
	var b uint16 = 20
	fmt.Printf("'a': %.3d - %.10b\n", a, a)
	fmt.Printf("'b': %.3d - %.10b\n\n", b, b)

	fmt.Printf("Bitwise AND: %d - %.10b\n", a&b, a&b)
	fmt.Printf("Bitwise OR: %d - %.10b\n", a|b, a|b)
	fmt.Printf("Bitwise XOR: %d - %.10b\n", a^b, a^b)
	fmt.Printf("Bitwise NOT: %d - %.10b\n", ^a, ^a)

	fmt.Println()
	fmt.Printf("Bitwise NAND: %d - %.10b\n", ^(a & b), ^(a & b))
	fmt.Printf("Bitwise NOR: %d - %.10b\n", ^(a | b), ^(a | b))
	fmt.Printf("Bitwise XNOR: %d - %.10b\n", ^(a ^ b), ^(a ^ b))

	READ_ROLE := 1        //0001
	WRITE_ROLE := 1 << 1  //0010
	UPDATE_ROLE := 1 << 2 //0100
	DELETE_ROLE := 1 << 3 //1000

	fmt.Println()
	fmt.Printf("Roles Binary: %.4b, %.4b, %.4b & %.4b \n", READ_ROLE, WRITE_ROLE, UPDATE_ROLE, DELETE_ROLE)
	fmt.Printf("Roles Number: %.4d, %.4d, %.4d & %.4d \n", READ_ROLE, WRITE_ROLE, UPDATE_ROLE, DELETE_ROLE)

	myProfile := READ_ROLE + WRITE_ROLE + DELETE_ROLE
	fmt.Printf("My profile: (%d, %b) \n", myProfile, myProfile)

	fmt.Println()

	fmt.Println("Does the user have permissions to perform the operation?")
	if (myProfile & DELETE_ROLE) != 0 {
		fmt.Println("Yes! :D")
	} else {
		fmt.Println("No :(")
	}
}
