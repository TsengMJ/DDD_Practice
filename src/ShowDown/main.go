package main

import (
	"showdown/showdown"
)

func main() {
	player_0 := &showdown.HumanPlayer{}
	player_1 := &showdown.HumanPlayer{}
	player_2 := &showdown.AIPlayer{}
	player_3 := &showdown.AIPlayer{}

	players := []showdown.IPlayer{player_0, player_1, player_2, player_3}

	deck := showdown.Deck{}
	deck.Init()

	showdown := &showdown.Showdown{}
	showdown.StartGame(players, deck)
}
