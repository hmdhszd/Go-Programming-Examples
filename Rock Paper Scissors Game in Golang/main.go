package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {

	playerChoice := ""
	PlayerValue := -1

	playerScore := 0
	computerScore := 0

	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= 3; i++ {

		clearScreen()

		fmt.Println("Rock, Paper & Scissors")
		fmt.Println("-----------------------")
		fmt.Println("Game is played for 3 rounds and the best of three wins the game. Good luck!")
		fmt.Println()

		fmt.Println("Round", i)
		fmt.Println("-------")
		fmt.Println()
		fmt.Print("Please enter rock, paper, scissors -> ")

		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		rand.Seed(time.Now().UnixNano())
		computerValue := rand.Intn(3)

		if playerChoice == "rock" {
			PlayerValue = ROCK
		} else if playerChoice == "paper" {
			PlayerValue = PAPER
		} else if playerChoice == "scissors" {
			PlayerValue = SCISSORS
		} else {
			PlayerValue = -1
		}

		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK !")
		case PAPER:
			fmt.Println("Computer chose PAPER !")
		case SCISSORS:
			fmt.Println("Computer chose SCISSORS !")
		default:
			fmt.Println("Default message. Computer chose NOTHING !")
		}

		fmt.Println("player chose", playerChoice)

		if computerValue == PlayerValue {

			fmt.Println("*** it's a draw ***")
			i--

		} else {

			switch PlayerValue {

			case ROCK:
				if computerValue == PAPER {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}

			case PAPER:
				if computerValue == SCISSORS {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}

			case SCISSORS:
				if computerValue == ROCK {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}

			default:
				fmt.Println("Invalid choise! please choose only : rock , paper or scissors")
				i--
			}

		}

		time.Sleep(5 * time.Second)

	}

	clearScreen()
	fmt.Println("Final Score:")
	fmt.Println("____________")
	fmt.Printf("Computer: %d/3, Player: %d/3", computerScore, playerScore)

	if computerScore > playerScore {
		fmt.Println("* * * * Computer Wins * * * * ")
	} else {
		fmt.Println("* * * * Player Wins * * * * ")
	}
}

func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {

		// Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

	} else {

		// Linux or Mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

	}
}

func computerWins(score int) int {
	fmt.Println("Computer Wins")
	return score + 1
}

func playerWins(score int) int {
	fmt.Println("Player Wins")
	return score + 1
}
