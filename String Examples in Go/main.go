package main

import (
	"fmt"
	"strings"
)

func main() {

	courses := []string{
		"I love Java programming language!",
		"I love Python programming language!",
		"I love Go programming language!",
		"I love Ruby programming language!",
	}

	for index, data := range courses {
		fmt.Println("index is:", index, "     data is:", data)
	}

	for index, data := range courses {
		if strings.Contains(data, "Go") {
			fmt.Println("\"Go\" is found in\"", data, "\" and index in :", strings.Index(data, "Go"), "is this string")
			fmt.Println("index of this string in the \"courses\" is :", index)
		}
	}

	//////////////////////////////////////////////////////////////////

	newString := "Go is a great programming language. Go for it!"

	for j := 0; j < len(newString); j++ {
		fmt.Print(string(newString[j]) + " ")
	}

	fmt.Println()
	fmt.Println("There are ", strings.Count(newString, "Go"), "\"Go\" in the newString")

	fmt.Println(strings.Index(newString, "G"))
	fmt.Println(strings.LastIndex(newString, "G"))

	//////////////////////////////////////////////////////////////////

	fmt.Println(newString)

	if strings.Contains(newString, "Go") {
		newString = strings.ReplaceAll(newString, "Go", "go")
	}

	fmt.Println(newString)

	if strings.Contains(newString, "go") {
		newString = strings.Replace(newString, "go", "Golang", 1)
	}

	fmt.Println(newString)

	//////////////////////////////////////////////////////////////////

	badEmail := "                   hamid@hamid.hamid                   "
	fmt.Printf("Bad Email is:\"%s\"", badEmail)

	fmt.Println()

	goodEmail := strings.TrimSpace(badEmail)
	fmt.Printf("Good Email is: \"%s\"", goodEmail)

	//////////////////////////////////////////////////////////////////

}
