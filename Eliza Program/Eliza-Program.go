package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	whatToSay := doctor.Intro()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(whatToSay)

	for {

		fmt.Print("-->")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {

			break

		} else {

			fmt.Println(doctor.Response(userInput))

		}

	}

}
