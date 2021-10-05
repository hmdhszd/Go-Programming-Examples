package main

import "fmt"

func main() {
	sayHelloWorld("Hello World fron inside of a function")
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
