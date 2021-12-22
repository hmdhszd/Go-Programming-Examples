package main

import (
	"fmt"
)

const (
	C1 = iota
	C2
	C3
)

const (
	D1 = iota + 1
	D2
	D3
)

const (
	x1 = iota + 1
	x2
	x3 = iota + 1
	x4
	x5
	x6
)

func main() {

	fmt.Println(C1, C2, C3)
	fmt.Println(D1, D2, D3)
	fmt.Println(x1, x2, x3, x4, x5, x6)

}
