package main

import (
	"fmt"
	"math"
)

func main() {

	answer := 7 + 3*4
	fmt.Println("answer is:", answer)

	answer = (7 + 3) * 4
	fmt.Println("answer is now:", answer)

	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("Area is:", area)

	half := 1 / 2
	fmt.Println("Half with integer division :", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("Half Float is:", halfFloat)

	badThreeSquared := 3 ^ 2
	fmt.Println("badThreeSquared is:", badThreeSquared)

	goodThreeSquared := math.Pow(3, 2)
	fmt.Println("goodThreeSquared is:", goodThreeSquared)

	remainder := 50 % 3
	fmt.Println("remainder is", remainder)

}
