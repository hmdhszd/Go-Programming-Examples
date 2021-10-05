package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)

	coffees[1] = "Cappucino"
	coffees[2] = "Americano"
	coffees[3] = "Latte"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"
	coffees[7] = "Quit"

	fmt.Println("   Menu")
	fmt.Println("----------")

	fmt.Println("1) Cappucino")
	fmt.Println("2) Americano")
	fmt.Println("3) Latte")
	fmt.Println("4) Mocha")
	fmt.Println("5) Macchiato")
	fmt.Println("6) Espresso")
	fmt.Println("Q) Quit")

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))

	}

	fmt.Println("Exiting Program . . .")
}
