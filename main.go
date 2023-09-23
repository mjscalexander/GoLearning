package main

import (
	"fmt"
  "encoding/json"
)

func main() {
  fmt.Println("Hello World")
  mstr, _ := json.Marshal("testing 123")
  err := json.Unmarshal(mstr)


}