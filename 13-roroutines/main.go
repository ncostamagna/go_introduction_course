package main

import (
	"fmt"
	"time"
)

func main() {

	go myProcess("A")
	go myProcess("B")

	time.Sleep(3 * time.Second)

	myFirstChannel := make(chan string)

	go myProcessWithChannel("C", myFirstChannel)

	result := <-myFirstChannel
	fmt.Println(result)
	close(myFirstChannel)

	channelA := make(chan string)
	channelB := make(chan string)

	go myProcessWithChannel("D", channelA)
	go myOtherProcessWithChannel("E", channelB)

	resultA := <-channelA
	fmt.Println(resultA)

	resultB := <-channelB
	fmt.Println(resultB)

	close(channelA)
	close(channelB)
}

func myProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 20 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "ok"
}

func myOtherProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 10 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}

	c <- "ok 2"
}

func myProcess(p string) {
	i := 0
	for i < 15 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}
