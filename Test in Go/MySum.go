package main

import "fmt"

func main() {
	fmt.Println("3 + 2 =", MySum(2, 3))
	fmt.Println("2 + 6 =", MySum(2, 6))
	fmt.Println("33 + 72 =", MySum(33, 72))
	fmt.Println("3453 + 2345 =", MySum(3453, 2345))
}

func MySum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}
	return sum
}
