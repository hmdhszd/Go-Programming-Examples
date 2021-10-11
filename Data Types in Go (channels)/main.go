package main

import "fmt"

func main() {

	go doSomething("Hello, world!")
	go doAnOtherThing("it does another thing!")

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

func doAnOtherThing(s string) {

	until := 0

	for {

		fmt.Println(s)
		until += 1

		if until >= 5 {
			break
		}

	}

}
