package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChannel chan rune

func main() {
	keyPressChannel = make(chan rune)

	go listenForKeyPress()

	fmt.Println("press any key or q to quit")

	_ = keyboard.Open()

	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()

		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChannel <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChannel
		fmt.Println("you pressed ", string(key))
	}

}
