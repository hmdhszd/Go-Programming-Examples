package main

import "fmt"

func main() {

	var FirstNumber int
	FirstNumber = 2

	var SecondNumber = 5

	ThirdNumber := 7

	var Answer int

	fmt.Println(FirstNumber)
	fmt.Println(SecondNumber)
	fmt.Println(ThirdNumber)
	fmt.Println(Answer)

	var MyFirstVariable string
	MyFirstVariable = "This is my First variable"

	sayHelloWorld(MyFirstVariable)

	MySecondVariable := "This is my Second variable"
	sayHelloWorld(MySecondVariable)

}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
