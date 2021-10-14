package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i <= 10; i++ {
		println("i is ->", i)
	}

	//===============================================================

	rand.Seed(time.Now().UnixNano())
	i := 1000

	for i > 100 {
		i = rand.Intn(1000) + 1

		fmt.Println("i is", i)

		if i > 100 {
			println("i is ", i, "So loop keeps going")
		}
	}

	fmt.Println("got", i, "Loop stopped")

}
