package main

import (
	"fmt"
  "slices"
)

func main() {
  lst := []string{"22", "1", "23","1", "25"}
  for _, num := range lst {
    flg, cache := uniqCache(num)
    if flg {
      fmt.Println("Id in list", cache)
    } else {
      fmt.Println("Id not in list, adding it...", cache)
    }
  }
}

func uniqCache(str string) (bool, []string) {
  newlist := []string{"21"}
  var flg bool
  flg = false
  func() {
    if slices.Contains(newlist,str) != true {
        newlist = append(newlist, str)
      } else {
        flg = true
      }
  }()
  return flg, newlist
}