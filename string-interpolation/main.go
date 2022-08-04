package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

var reader *bufio.Reader

func main() {
  reader = bufio.NewReader(os.Stdin)
  userName := readString("What is your name?")
  age := readInt("How old are you?")
  
  
  // fmt.Println("Your name is" + userName + ". You are", age, "years old.")
  //fmt.Println(fmt.Sprintf("Your name is %s. you are %d years old", userName, age))
  fmt.Printf("Your name is %s. you are %d years old.\n", userName, age)
}

func prompt() {
  fmt.Print("-> ")
}
//  func param s == string;
func readString(s string) string { // --> output == string
  for {
  fmt.Println(s)
  prompt()

  userInput, _ := reader.ReadString('\n')

  userInput = strings.Replace(userInput, "\r\n", "", -1)
  userInput = strings.Replace(userInput, "\n", "", -1)
  if userInput == "" {
    fmt.Println("Please enter a string value")
  } else {
      return userInput
  }
  }
}


func readInt(s string) int {
  for {
  fmt.Println(s)
  prompt()

  userInput, _ := reader.ReadString('\n')
  userInput = strings.Replace(userInput, "\r\n", "", -1)
  userInput = strings.Replace(userInput, "\n", "", -1)

  num, err := strconv.Atoi(userInput)
  if err != nil {
    fmt.Println("Please enter a whole number")
  } else {
      return num
  }

  }
}