package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	if err := godotenv.Load("otherFolder/.var"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	myEnv, err := godotenv.Read("otherFolder/.var")
	fmt.Println(err)
	fmt.Println(myEnv)

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

}
