package main

import (
  "bufio"
  "fmt"
  "math/rand"
  "os"
  "time"
)

const prompt = "and don't type your number in, just press ENTER when ready"

func main() {
  // seed the random number generator
  rand.Seed(time.Now().UnixNano())

  // rand generate a number between 0 and whatever is passed as a parameter
  // we add 2 to it because we want the number used to be at least 2 and no greater than 10 (multiplying by 1 makes the game a bit silly)
  var firstNumber = rand.Intn(8) + 2
  var secondNumber = rand.Intn(8) + 2
  var subtraction =  rand.Intn(8) + 2
  var answer = firstNumber * secondNumber - subtraction

 game_logic(firstNumber, secondNumber, subtraction, answer)
}

func game_logic(firstNumber, secondNumber, subtraction, answer int) {
  
  // create our reader variable, which reads input from stdin (the keyboard)
  reader := bufio.NewReader(os.Stdin)
  
  fmt.Println("Guess the Number Game")
  fmt.Println("---------------------")
  fmt.Println("")

  fmt.Println("Think of number between 1 and 10", prompt)
  reader.ReadString('\n')
  
  fmt.Println("Multiply you number by", firstNumber, prompt)
  reader.ReadString('\n')
  
  fmt.Println("Now multiply the result by", secondNumber, prompt)
  reader.ReadString('\n')
  
  fmt.Println("Divide the result by number you originally thought of", prompt)
  reader.ReadString('\n')

  fmt.Println("Now subtract", subtraction, prompt)
  reader.ReadString('\n')

  answer = firstNumber * secondNumber - subtraction
  fmt.Println("the answer is", answer)
}