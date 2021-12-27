package main

import "fmt"

func main() {

	// size of the channel is 2
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

}
