package main

import (
	"fmt"
)

func main() {
  
  
  
  uniqCache()

  



}

func uniqCache() {
  var newlst []string

  func() {
    newlst = append(newlst, "34")
    fmt.Println(newlst)

  }()
}