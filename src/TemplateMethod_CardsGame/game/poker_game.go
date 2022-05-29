package game

import (
	"fmt"
	"template_method/game/deck"
	"template_method/game/player"
)

type PokerGame struct {
	CardGame
}

func (pg *PokerGame) Playing(players []player.IPlayer) {
	pg.setPlayers(players)
	pg.createDeck()
	pg.start()
	pg.end()
}

func (pg *PokerGame) createDeck() {
	cardTypeKeys := make([]string, 0)
	cardValueKeys := make([]string, 0)

	cards := make([]deck.PokerCard, 0)
	for cardTypeKey, _ := range deck.PokerType {
		cardTypeKeys = append(cardTypeKeys, cardTypeKey)
	}

	for cardValueKey, _ := range deck.PokerValue {
		cardValueKeys = append(cardValueKeys, cardValueKey)
	}

	for _, cardTypeKey := range cardTypeKeys {
		for _, cardValueKey := range cardValueKeys {
			var newCard = deck.PokerCard{Type: cardTypeKey, Value: cardValueKey}
			cards = append(cards, newCard)
		}
	}

	newDeck := deck.Deck{Cards: cards, Next: 0}

	pg.deck = newDeck
	pg.deck.Shuffle()

	for i := 0; i < len(pg.deck.Cards); i++ {
		pg.palyers[i%4].DrawCard(&pg.deck)
	}
}

func (pg *PokerGame) start() {
	for pg.palyers[0].NumberOfHands() > 0 {
		pg.nextRound()
	}
}

func (pg *PokerGame) nextRound() {

	showCards := make([]deck.Card, 0)

	for i := 0; i < len(pg.palyers); i++ {
		showCard := pg.palyers[i].ShowCard()
		showCards = append(showCards, showCard)
	}

	maxCard := showCards[0]
	maxIndex := 0
	for i := 1; i < len(showCards); i++ {
		fmt.Println(showCards[i].Compare(maxCard))
		if showCards[i].Compare(maxCard) > 0 {
			maxCard = showCards[i]
			maxIndex = i
		}
	}

	pg.palyers[maxIndex].AddPoint()
}

func (pg *PokerGame) end() {
	fmt.Println("============= FINAL SCORE  =============")

	maxPoint := 0
	maxIndex := 0
	for i := 0; i < len(pg.palyers); i++ {
		fmt.Printf("Player %d: %d \n", i, pg.palyers[i].GetPoint())
		if pg.palyers[i].GetPoint() > maxPoint {
			maxPoint = pg.palyers[i].GetPoint()
			maxIndex = i
		}
	}

	fmt.Printf("The Winner is: Player %d\n", pg.palyers[maxIndex].GetPoint())
}
