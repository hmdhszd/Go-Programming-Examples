package main

import "fmt"

func main() {

	go doSomething("Hello, world!")

	fmt.Println("this is another message!")

	for {
		// do nothing
	}

}

func doSomething(s string) {

	until := 0

	for {

		fmt.Println(s)
		until += 1

		if until >= 5 {
			break
		}

	}

}
