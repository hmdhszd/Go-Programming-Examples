package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("a %s says %s . ", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("a %s has %d legs. ", a.Name, a.NumberOfLegs)
}

func main() {


	//-------------------------------------------------
	//-------------------------------------------------

	ii := []int{3,5,2,6,7,2,4,7,3,3,7}

	s := sum(ii...)

	fmt.Println(s)

	//-------------------------------------------------
	//-------------------------------------------------

	z := addTwoNumbers(2, 4)
	fmt.Println(z)

	//-------------------------------------------------

	myTotal := sumManyNumbers(2, 3, 4, 5, -4)
	println(myTotal)

	//-------------------------------------------------

	var dog Animal
	dog.Name = "Dog"
	dog.Sound = "Hap Hap"
	dog.NumberOfLegs = 4

	dog.Says()
	dog.HowManyLegs()

	//-------------------------------------------------

	cat := Animal{
		Name:         "cat",
		Sound:        "Mioooow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()

}

func addTwoNumbers(x, y int) int {
	return x + y
}

func sumManyNumbers(nums ...int) int {

	total := 0

	for _, x := range nums {
		total += x
	}

	return total
}


func sum(x ...int) int {

	total := 0

	for _,v := range x {
		total += v
	}

	return total
	}