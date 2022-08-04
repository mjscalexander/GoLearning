package main

import (
  "myapp/packageone"
)

var myVar = "hello"

func main() {
  var blockVar = "world" 
  packageone.PackageVar = "from package one"
  packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
  
}