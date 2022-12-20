package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

// This function (main) it's the play tennis game
// You need to pass a player one and player two names
// This function end when match is over
// The match is over when one play won a game
func main() {
	var playerOne, playerTwo string

	fmt.Print("Input player name one: ")
	_, err1 := fmt.Scanln(&playerOne)

	fmt.Print("Input player name two: ")
	_, err2 := fmt.Scanln(&playerTwo)

	if err1 != nil || len(playerOne) == 0 {
		log.Fatal("Unable to read player name one!")
	} else if err2 != nil || len(playerTwo) == 0 {
		log.Fatal("Unable to read player name two!")
	}

	waitGroup.Add(1)

	go playTennis(playerOne, playerTwo)
	waitGroup.Wait()
}

// This function simulates a tennis match.
// The match starts with the players with the score of 0 to 0.
// If the player fails to throw the ball back, a point is scored.
func playTennis(playerOne string, playerTwo string) {
	var score = make(map[string]int)
	score[playerOne] = 0
	score[playerTwo] = 0

	ball := make(chan bool)
	playerKicking := playerOne

	for true {
		winner := playerWinner(score, playerOne, playerTwo)

		if len(winner) > 0 {
			fmt.Println("The winner is", winner, "!!!")
			break
		}

		go kickBack(ball)

		playAgain := <-ball

		if !playAgain {
			score[playerKicking] += 1
		} else {
			if playerKicking == playerOne {
				playerKicking = playerTwo
			} else {
				playerKicking = playerOne
			}
		}

		fmt.Println("Score", playerOne, score[playerOne], "and", playerTwo, score[playerTwo])
	}

	waitGroup.Done()
}

// This function simulate a kickback the ball
func kickBack(ball chan bool) {
	rand.Seed(time.Now().UnixNano())
	ball <- rand.Intn(2) == 1
}

// This function determines who won the match.
func playerWinner(score map[string]int, playerOne string, playerTwo string) string {
	scorePlayerOne := score[playerOne]
	scorePlayerTwo := score[playerTwo]

	if winnerInTime(scorePlayerOne, scorePlayerTwo) {
		return playerOne
	} else if winnerInDeuce(scorePlayerOne, scorePlayerTwo) {
		return playerOne
	} else if winnerInTime(scorePlayerTwo, scorePlayerOne) {
		return playerTwo
	} else if winnerInDeuce(scorePlayerTwo, scorePlayerOne) {
		return playerTwo
	}

	return ""
}

// This function determines who won the match up to the fourth point.
func winnerInTime(scoreOne int, scoreTwo int) bool {
	return (scoreOne == 4) && (scoreTwo == 0 || scoreOne-scoreTwo > 2 || scoreOne-scoreTwo == 2)
}

// This function determines who wins after a 3-3 tie.
func winnerInDeuce(scoreOne int, scoreTwo int) bool {
	return (scoreOne > 4) && (scoreOne-scoreTwo == 2)
}
