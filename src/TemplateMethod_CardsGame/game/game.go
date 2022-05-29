package game

import (
	"fmt"
	"template_method/game/deck"
	"template_method/game/player"
)

type IGame interface {
	PlayingGame(players []player.IPlayer)
	initDeck()
	initPlayers(players []player.IPlayer)
	startGame()
	playing()
	endGame()
}

type Game struct {
	IGame
	Players []player.IPlayer
	Deck    deck.IDeck
}

func (g *Game) PlayingGame(players []player.IPlayer) {
	g.initDeck()
	g.initPlayers(players)
	g.startGame()
	g.playing()
	g.endGame()
}

func (g *Game) initDeck() {
	g.Deck = &deck.Deck{}
}

func (g *Game) initPlayers(players []player.IPlayer) {
	g.Players = players
	for i := 0; i < len(g.Players); i++ {
		fmt.Printf("-- Player %d --\n", i)
		g.Players[i].ReName()
		fmt.Printf("Hello %s\n\n", g.Players[i].GetName())
	}
}

func (g *Game) startGame() {
	fmt.Println("Game Start")
	g.playing()
}

func (g *Game) playing() {
	fmt.Println("Playing Game")
}

func (g *Game) endGame() {
	fmt.Println("Game End")
}
