package main

import (
	"fmt"
	"math/rand"
)

func main() {
	channel := make(chan string)
	numRounds := 3
	
	go throwingNinjaStart(channel, numRounds)
	
	for i := 0 ; i < numRounds; i++ {
		fmt.Println(<-channel)
	}
	
}

func throwingNinjaStart(channel chan string, numRounds int) {
	for i := 0 ; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}
}