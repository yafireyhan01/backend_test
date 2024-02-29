package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	id     int
	dice   []int
	points int
}

// Permainan Dadu-------

func main() {
	var numPlayers, numDice int

	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scanln(&numPlayers)

	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scanln(&numDice)

	players := make([]Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = Player{id: i + 1, dice: kocokDadu(numDice)}
	}
	
	turn := 1
	for {
		fmt.Printf("Turn %d\n\n", turn)
		for i, player := range players {
			fmt.Printf("Player #%d (%d): %v\n", player.id, player.points, player.dice)
			evaluasiDadu(&players[i], players)
		}

		activePlayers := 0
		for _, player := range players {
			if len(player.dice) > 0 {
				activePlayers++
			}
		}
		if activePlayers <= 1 {
			break
		}

		fmt.Println("==================")
		turn++
	}

	maxPoints := 0
	winner := 0
	for i, player := range players {
		if player.points > maxPoints {
			maxPoints = player.points
			winner = i + 1
		}
	}

	fmt.Printf("Game dimenangkan oleh pemain #%d  karena memiliki poin lebih banyak dari pemain lainnya.\n", winner)
}

func kocokDadu(numDice int) []int {
	dice := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		dice[i] = rand.Intn(6) + 1
	}
	return dice
}

func evaluasiDadu(player *Player, players []Player) {
	var newDice []int
	for _, d := range player.dice {
		switch d {
		case 1:
			if player.id == len(players) {
				players[0].dice = append(players[0].dice, 1)
			} else {
				players[player.id].dice = append(players[player.id].dice, 1)
			}
		case 6:
			player.points++
		default:
			newDice = append(newDice, d)
		}
	}
	player.dice = newDice
}
