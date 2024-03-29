package main

import (
	"fmt"
)

func main() {
	ninja1, ninja2 := make(chan string), make(chan string)

	go captainElect(ninja1, "Ninja1")
	go captainElect(ninja2, "Ninja2")

	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	default:
		fmt.Println("Neither was chosen")
	}
	roughlyFair()
}

func captainElect(ninja chan string, message string) {
	
	ninja <- message
}

func roughlyFair() {
	ninja1 := make(chan interface{}); close(ninja1)
	ninja2 := make(chan interface{}); close(ninja2)

	var ninja1Count, ninja2Count int

	for i:=0; i < 1000; i++ {
		select {
		case <- ninja1:
			ninja1Count++
		case <- ninja2:
			ninja2Count++
		}
	}
	fmt.Printf("ninja1Count: %d, ninja2Count: %d\n", ninja1Count,ninja2Count)
}