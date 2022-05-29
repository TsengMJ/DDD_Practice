package player

import (
	"fmt"
	"math/rand"
	"template_method/game/deck"
)

type AIPlayer struct {
	Player
}

func (p *AIPlayer) ShowCard() deck.ICard {
	cardIndex := rand.Intn(p.GetNumberOfCards())
	showCard := p.Cards[cardIndex]
	p.removeCard(cardIndex)

	return showCard
}
