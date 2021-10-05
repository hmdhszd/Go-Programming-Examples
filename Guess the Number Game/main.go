package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const MySentence = "then press Enter when you are ready!"

func main() {

	rand.Seed(time.Now().UnixNano())

	FirstNumber := rand.Intn(8) + 2
	SecondNumber := rand.Intn(8) + 2
	Subtraction := rand.Intn(8) + 2

	Answer := FirstNumber*SecondNumber - Subtraction

	PlayTheGame(FirstNumber, SecondNumber, Subtraction, Answer)

}

func PlayTheGame(FirstNumber, SecondNumber, Subtraction, Answer int) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number Game")
	fmt.Println("_-_-_-_-_-_-_-_-_-_-_-_")
	fmt.Println()

	fmt.Println("Think of a number between 1 and 10", MySentence)

	reader.ReadString('\n')

	fmt.Println("Multyply your number by", FirstNumber, MySentence)

	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", SecondNumber, MySentence)

	reader.ReadString('\n')

	fmt.Println("Devide the result by the number you originaly thought of", MySentence)

	reader.ReadString('\n')

	fmt.Println("Now substract", Subtraction, MySentence)

	reader.ReadString('\n')

	fmt.Println("Th answer is ", Answer)

}
