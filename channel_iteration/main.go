package main

import (
	"fmt"
	"math/rand"
)

func main() {
	channel := make(chan string)	
	go throwingNinjaStart(channel)
	
	for msg := range channel {
		fmt.Println(msg)
	}

}

func throwingNinjaStart(channel chan string) {
	numRounds := 3
	for i:= 0 ; i < numRounds ; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}
	close(channel)
	// can't reopen a closed channel. you have to create a new one
}