package main

import "fmt"

type Animal struct {
	Name string
	Sound string
	NumberOfLegs int
}

func(a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name , a.Sound)
	fmt.Println()
}

func(a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d of Legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()

	cat := Animal {
		Name: "cat",
		Sound: "meow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()
}

func sumMany(nums ...int) int {
	total := 0

	for _,x := range nums {
		total = total + x
	}

	return total
}
