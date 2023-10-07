package main

import (
	"fmt"
)

func main() {
	channel := make(chan string,2)
	channel <- "New Message"
	channel <- "Second Message"
	// go func(){
	// 	channel <- "First Message"
	// }()
	
	fmt.Println(<-channel)
}

