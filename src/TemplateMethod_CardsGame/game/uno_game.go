package game

import (
	"fmt"
	"template_method/game/deck"
	"template_method/game/player"
)

type UnoGame struct {
	Game
}

func (g *UnoGame) PlayingGame(players []player.IPlayer) {
	g.initDeck()
	g.initPlayers(players)
	g.startGame()
	g.playing()
	g.endGame()
}

func (g *UnoGame) initDeck() {
	g.Deck = &deck.Deck{}
}

func (g *UnoGame) startGame() {
	fmt.Println("UNO Game Start")
	g.playing()
}

func (g *UnoGame) playing() {
	fmt.Println("UNO Poker Game")
}

func (g *UnoGame) endGame() {
	fmt.Println("UNO Game End")
}
