package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const p1 = "Frank"
const p2 = "Ohanna"

var waitGroup sync.WaitGroup

/*
1 - Um game é uma partida que ganha o adversário que marcar pelo menos 4 pontos no total e 2 a mais que o adversário
2 - Um set é um conjunto de jogos, vencendo o jogador que ganhar pelo menos 6 jogos e dois a mais que o adversário
3 - Uma partida (match) é uma sequência de sets que ganha 3 de 5 sets

Simplicidade: Um match possui apenas um único set, composto de um único game, ganhando o jogador que fizer um número P de pontos
*/
func playTennis(p1 string, p2 string) {
	var score = make(map[string]int)
	score[p1] = 0
	score[p2] = 0

	ball := make(chan bool)
	playerKicking := p1

	for true {
		playerWinner := winner(score, p1, p2)

		if len(playerWinner) > 0 {
			fmt.Println("The winner is", playerWinner,"!!!")
			break
		}

		go kickBack(ball)

		playAgain := <-ball

		if !playAgain {
			score[playerKicking] += 1
		} else {
			if playerKicking == p1 {
				playerKicking = p2
			} else {
				playerKicking = p1
			}
		}

		fmt.Println("Score", p1, score[p1], "and", p2, score[p2])
	}

	waitGroup.Done()
}

func kickBack(ball chan bool) {
	rand.Seed(time.Now().UnixNano())
	ball <- rand.Intn(2) == 1
}

func winner(score map[string]int, p1 string, p2 string) string {
	if score[p1] == 4 && (score[p2] == 0 || score[p1]-score[p2] > 2 || score[p1]-score[p2] == 2) {
		return p1
	} else if score[p1] > 4 && score[p1]-score[p2] == 2 {
		return p1
	} else if score[p2] == 4 && (score[p1] == 0 || score[p2]-score[p1] > 2 || score[p1]-score[p2] == 2)  {
		return p2
	} else if score[p2] > 4 && score[p2]-score[p1] == 2 {
		return p2
	} else {
		return ""
	}
}

func main() {
	waitGroup.Add(1)
	go playTennis(p1, p2)
	waitGroup.Wait()
}
