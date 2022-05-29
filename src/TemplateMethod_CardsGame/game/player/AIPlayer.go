package player

import (
	"math/rand"
	"template_method/game/deck"
)

type AIPlayer struct {
	Player
}

func (p *AIPlayer) ShowCard() deck.Card {
	cardIndex := rand.Intn(len(p.Hands))
	showCard := p.Hands[cardIndex]
	p.RemoveCard(cardIndex)

	return showCard
}
