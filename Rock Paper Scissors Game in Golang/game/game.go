package game

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

type Game struct {
	DisplayChannel chan string
	RoundChannel   chan int
	Round          Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {

	for {

		select {

		case round := <-g.RoundChannel:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChannel <- 1

		case msg := <-g.DisplayChannel:
			fmt.Println(msg)
		}

	}
}

func (g *Game) ClearScreen() {
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

func (g *Game) PrintIntro() {

	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("-----------------------")
	fmt.Println("Game is played for 3 rounds and the best of three wins the game. Good luck!")
	fmt.Println()

}

func (g *Game) PlayRound() bool {

	rand.Seed(time.Now().UnixNano())

	playerValue := -1

	fmt.Println()
	fmt.Println("Round", g.Round.RoundNumber)
	fmt.Println("--------")
	fmt.Print("Please enter rock, paper, scissors -> ")

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()
	g.DisplayChannel <- fmt.Sprintf("player chose %s", strings.ToUpper(playerChoice))

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

	if computerValue == playerValue {
		g.DisplayChannel <- "*** it's a draw ***"
		return false

	} else {

		switch playerValue {

		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}

		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}

		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}

		default:
			g.DisplayChannel <- "Invalid choise! please choose only : rock , paper or scissors"
			return false

		}

	}

	return true

}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChannel <- "Computer Wins!"
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChannel <- "Player Wins!"
}

func (g *Game) PrintSummary() {

	fmt.Println("Final Score:")
	fmt.Println("____________")
	fmt.Println()
	fmt.Printf("Computer: %d/3, Player: %d/3", g.Round.ComputerScore, g.Round.PlayerScore)

	if g.Round.ComputerScore > g.Round.PlayerScore {
		fmt.Println()
		fmt.Println("* * * * Computer Wins * * * * ")
	} else {
		fmt.Println("* * * * Player Wins * * * * ")
	}

}
