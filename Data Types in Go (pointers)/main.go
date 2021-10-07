package main

import "fmt"

func main() {

	x := 10
	myFirstPointer := &x

	fmt.Println("x is ", x)
	fmt.Println("myFirstPointer is ", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println("x is changed to ", x)

	changeValueOfPointer(&x)
	//changeValueOfPointer(myFirstPointer)
	fmt.Println("x after changeValueOfPointer function call: ", x)

}

func changeValueOfPointer(num *int) {
	*num = 25
}
