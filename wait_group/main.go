package main

import (
	"fmt"
	"log"
	"net/http"
)

var urls = []string{
	"https://twitter.com",
	"https://tutorialedge.net",
	"https://google.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	for _, url := range urls {
		
	}
}

func main(){
	fmt.Println("Go WaitGroup Tutorial")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}