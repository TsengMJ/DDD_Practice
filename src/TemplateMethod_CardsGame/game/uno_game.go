package game

import (
	"fmt"
	"template_method/game/deck"
	"template_method/game/player"
	"template_method/utils"
)

type UnoGame struct {
	Game
	poolCards []deck.ICard
}

func (g *UnoGame) PlayingGame(players []player.IPlayer) {
	g.initDeck()
	g.initPlayers(players)
	g.startGame()
	g.playing()
	g.endGame()
}

func (g *UnoGame) initDeck() {
	cardTypeKeys := utils.GetMapKeys(deck.UnoType)
	cardValueKeys := utils.GetMapKeys(deck.UnoValue)

	cards := make([]deck.ICard, 0)
	for _, cardTypeKey := range cardTypeKeys {
		for _, cardValueKey := range cardValueKeys {
			newCard := deck.PokerCard{}
			newCard.SetType(cardTypeKey)
			newCard.SetValue(cardValueKey)

			cards = append(cards, deck.ICard(&newCard))
		}
	}

	var newDeck = deck.Deck{Cards: cards}
	g.Deck = deck.IDeck(&newDeck)
}

func (g *UnoGame) startGame() {
	fmt.Printf("\n-------------------\n")
	fmt.Println("| UNO Game Start |")
	fmt.Printf("-------------------\n")

	// DrawCards
	initialCardsPerPlayer := 5
	for i := 0; i < initialCardsPerPlayer*len(g.Players); i++ {
		g.Players[i%4].DrawCards(g.Deck)
	}
}

func (g *UnoGame) playing() {

	fmt.Println("P - 1")
	g.checkDeck()
	targetCard := g.Deck.Draw()

	for i := 0; i < len(g.Players); i++ {
		showCard := g.Players[i].ShowCard()

		if showCard.Compare(targetCard) > 0 {
			targetCard = showCard
		} else {
			fmt.Println("P - 2")
			g.checkDeck()
			g.Players[i].DrawCards(g.Deck)
			g.Players[i].AddCard(targetCard)
		}
	}

	nextRound := true
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i].GetNumberOfCards() == 0 {
			nextRound = false
		}
	}

	if nextRound {
		g.playing()
	}
}

func (g *UnoGame) endGame() {
	fmt.Printf("\n-------------------\n")
	fmt.Println("| UNO Game End |")
	fmt.Printf("-------------------\n")

	// Show all player's left Cards
	winner := g.Players[0]
	leftCards := winner.GetNumberOfCards()
	for i := 0; i < len(g.Players); i++ {
		fmt.Printf("\nPlayer %s left %d cards\n", g.Players[i].GetName(), g.Players[i].GetNumberOfCards())

		if g.Players[i].GetNumberOfCards() < leftCards {
			winner = g.Players[i]
			leftCards = winner.GetNumberOfCards()
		}
	}

	// Show the Winner and his score
	fmt.Printf("\nThe Winner is %s with %d cards left\n", winner.GetName(), leftCards)
}

func (g *UnoGame) checkDeck() {
	if len(g.Deck.GetCards()) == 0 {
		g.Deck.SetCards(g.poolCards)
		g.Deck.Shuffle()
	}
}
