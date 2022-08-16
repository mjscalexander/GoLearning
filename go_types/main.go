package main

import "fmt"

type Car struct {
	NumberOfTires  int
	Luxury         bool
	BucketSeats    bool
	Make           string
	Model          string
	Year           int
}


func main(){
	myCar := Car{
		NumberOfTires: 4,
		Luxury: true,
		Make: "Volvo",
		Model: "XC90",
		Year: 2019,
	}
	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)

	var myInt int

	myInt = 10

	fmt.Println(myInt)

	x := 10
	// hold the address of where "x" is
	myFirstPointer := &x
	
	fmt.Println("x is ", x)
	fmt.Println("my first pointer is ", myFirstPointer)
	
	// the '*' before the variable allow you to access the value at 
	// the actual address and change it if you like.
	*myFirstPointer = 15

	changeValueOfPointer(&x)

	fmt.Println("after function call, x is now", x)
}	

func changeValueOfPointer(num *int) {
	*num = 25
}