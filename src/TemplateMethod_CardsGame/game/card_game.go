package game

import (
	"fmt"
	"template_method/game/deck"
	"template_method/game/player"
)

type ICardGane interface {
	Playing(players []player.IPlayer)
	setPlayers(players []player.IPlayer)
	createDeck()
	start()
	end()
}

type CardGame struct {
	ICardGane
	palyers []player.IPlayer
	deck    deck.Deck
	isEnd   bool
}

func (cd *CardGame) Playing(players []player.IPlayer) {
	cd.setPlayers(players)
	cd.createDeck()
	cd.nextRound()
}

func (cd *CardGame) createDeck() {}

func (cd *CardGame) setPlayers(players []player.IPlayer) {
	cd.palyers = players

	for i := 0; i < len(cd.palyers); i++ {
		cd.palyers[i].SetOrder(i)
		cd.palyers[i].Rename()
	}
}

func (cd *CardGame) start() {}

func (cd *CardGame) nextRound() {
	if cd.isEnd {
		cd.end()
	} else {
		cd.isEnd = true
		cd.nextRound()
	}
}

func (cd *CardGame) end() {
	fmt.Println("End of the game")
}
