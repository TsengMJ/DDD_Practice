package player

import (
	"math/rand"
	"template_method/game/deck"
)

type AIPlayer struct {
	Player
}

func (p *AIPlayer) ShowCard() deck.ICard {
	cardIndex := rand.Intn(len(p.Cards))
	showCard := p.Cards[cardIndex]
	p.removeCard(cardIndex)

	return showCard
}
