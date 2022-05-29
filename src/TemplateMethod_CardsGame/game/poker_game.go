package game

import (
	"fmt"
	"template_method/game/deck"
	"template_method/game/player"
	"template_method/utils"
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
	cardTypeKeys := utils.GetMapKeys(deck.PokerType)
	cardValueKeys := utils.GetMapKeys(deck.PokerValue)

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

func (g *PokerGame) startGame() {
	fmt.Printf("\n-------------------\n")
	fmt.Println("| Poker Game Start |")
	fmt.Printf("-------------------\n")

	// DrawCards
	for i := 0; len(g.Deck.GetCards()) > 0; i++ {
		g.Players[i%4].DrawCards(g.Deck)
	}
}

func (g *PokerGame) playing() {
	// Every player show one card
	cards := make([]deck.ICard, 0)
	for i := 0; i < len(g.Players); i++ {
		cards = append(cards, g.Players[i].ShowCard())
	}

	// Select the winner
	maxCard := cards[0]
	maxIndex := 0
	for i := 0; i < len(cards); i++ {
		if cards[i].Compare(maxCard) > 0 {
			maxCard = cards[i]
			maxIndex = i
		}
	}

	g.Players[maxIndex].AddPoint()

	if g.Players[0].GetNumberOfCards() > 0 {
		fmt.Println("Keep Playing")
		g.playing()
	}
}

func (g *PokerGame) endGame() {
	fmt.Printf("\n-------------------\n")
	fmt.Println("| Poker Game End |")
	fmt.Printf("-------------------\n")

	// Show all player's score
	winner := g.Players[0]
	score := winner.GetPoint()
	for i := 0; i < len(g.Players); i++ {
		fmt.Printf("\nPlayer %s get %d scores\n", g.Players[i].GetName(), g.Players[i].GetPoint())

		if g.Players[i].GetPoint() > score {
			winner = g.Players[i]
			score = winner.GetPoint()
		}
	}

	// Show the Winner and his score
	fmt.Printf("\nThe Winner is %s with %d scores\n", winner.GetName(), score)
}
