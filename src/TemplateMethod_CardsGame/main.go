package main

import (
	"template_method/game"
	"template_method/game/player"
)

func main() {
	player_0 := &player.AIPlayer{}
	player_1 := &player.AIPlayer{}
	player_2 := &player.AIPlayer{}
	player_3 := &player.AIPlayer{}

	players := []player.IPlayer{player_0, player_1, player_2, player_3}

	pokerGame := &game.PokerGame{}

	pokerGame.Playing(players)
}
