package packageone

import "fmt"

var PackageVar string

func PrintMe(myVar, blockVar, PackageVar string) {
  fmt.Println(myVar, blockVar, PackageVar)
}