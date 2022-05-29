package main

import (
	"template_method/game"
	"template_method/game/player"
)

func main() {
	// Create Players
	player_0 := &player.AIPlayer{}
	player_1 := &player.AIPlayer{}
	player_2 := &player.AIPlayer{}
	player_3 := &player.AIPlayer{}

	players := []player.IPlayer{player_0, player_1, player_2, player_3}

	// Poker Game
	pokerGame := &game.PokerGame{}
	pokerGame.PlayingGame(players)

	// Uno Game
	unoGame := &game.UnoGame{}
	unoGame.PlayingGame(players)
}
