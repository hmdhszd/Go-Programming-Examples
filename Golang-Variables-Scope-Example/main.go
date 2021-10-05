package main

import (
	"fmt"
	"log"

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

	fmt.Println("Press any key on keyboard. Pres ESC to exit.")

	for {
		char, key, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("You pressed Char: ", char, "Key: ", key)
		} else {
			fmt.Println("You pressed Char: ", char, "Key: ", key)
		}

		if key == keyboard.KeyEsc {
			break
		}

	}

	fmt.Println("Exiting Program . . .")
}
