package main

import (
	"fmt"
  "slices"
)

func main() {  
  randint := "22"
  flg, lst := uniqCache(randint)
  if flg {
    fmt.Println("Id in list", lst)
  } else {
    fmt.Println("Id not in list, adding it...", lst)
  }

}

func uniqCache(str string) (bool, []string) {
  newlist := []string{"21"}
  var flg bool
  flg = false
  func() {
    if slices.Contains(newlist,str) == true {
      flg = true
      } else {
        newlist = append(newlist, str)
      }
  }()

  return flg, newlist
}