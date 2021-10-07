package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

type User struct {
	UserName       string
	Age            int64
	FavoriteNumber float64
	OwnsAdog       bool
}

var reader *bufio.Reader

func main() {
	var user User

	reader = bufio.NewReader(os.Stdin)

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")
	user.OwnsAdog = readBool("Do you own a dog? (y/n)")

	fmt.Printf("Your name is %s and you are %d years old. Your favorite number is %f . Own a dog: %t", user.UserName, user.Age, user.FavoriteNumber, user.OwnsAdog)

}

func prompt() {
	fmt.Print("--> ")
}

func readString(s string) string {

	for {
		fmt.Println(s)

		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput != "" {
			return userInput
		} else {
			fmt.Println("Please enter a value.")
		}
	}
}

func readInt(s string) int64 {

	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole numbers !")

		} else {
			return int64(num)
		}

	}
}

func readFloat(s string) float64 {

	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a numbers !")

		} else {
			return num
		}

	}
}

func readBool(s string) bool {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)

		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please Enter y or n")

		} else if char == 'n' || char == 'N' {
			return false

		} else if char == 'y' || char == 'Y' {
			return true
		}
	}

}
