package main

import (
	"myapp/game"
)

func main() {

	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{

		DisplayChannel: displayChan,
		RoundChannel:   roundChan,

		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	go game.Rounds()
	game.ClearScreen()
	game.PrintIntro()

	for {
		game.RoundChannel <- 1
		<-game.RoundChannel

		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChannel <- -1
			<-game.RoundChannel
		}
	}

	game.ClearScreen()
	game.PrintSummary()

}
