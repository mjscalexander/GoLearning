package main

import (
	"fmt"
	"sync"
	"net/http"
	"log"
)
// Simple example
// func myFunc(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Finished Executing Goroutine")
// 	wg.Done()
// }

// func main() {
// 	fmt.Println("Golang WaitGroup Tutorial")
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go myFunc(&wg)
// 	wg.Wait() // Will Block until wg reaches 0
// 	fmt.Println("Finished executing my go program")
// }

var urls = []string{
	"https://twitter.com",
	"https://www.openpolicyagent.org/",
	"https://wix.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) { 
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w,"Error: %+v\n",err)
			}
			if resp.StatusCode <= 200 {
				fmt.Fprintf(w, "%+v\n", resp.Status)
			}
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main(){
	fmt.Println("Go WaitGroup Tutorial")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}