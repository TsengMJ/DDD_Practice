package game

import (
	"fmt"
	"template_method/game/deck"
	"template_method/game/player"
)

type PokerGame struct {
	Game
}

func (g *PokerGame) PlayingGame(players []player.IPlayer) {
	g.initDeck()
	g.initPlayers(players)
	g.startGame()
	g.playing()
	g.endGame()
}

func (g *PokerGame) initDeck() {
	g.Deck = &deck.Deck{}
}

func (g *PokerGame) startGame() {
	fmt.Println("Poker Game Start")
	g.playing()
}

func (g *PokerGame) playing() {
	fmt.Println("Playing Poker Game")
}

func (g *PokerGame) endGame() {
	fmt.Println("Poker Game End")
}
