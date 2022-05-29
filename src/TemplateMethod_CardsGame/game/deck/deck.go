package deck

import (
	"math/rand"
	"time"
)

type IDeck interface {
	Shuffle()
	Draw() ICard
	SetCards(cards []ICard)
	GetCards() []ICard
}

type Deck struct {
	IDeck
	Cards []ICard
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) Draw() ICard {
	nextCard := d.Cards[0]
	d.Cards = d.Cards[1:]
	return nextCard
}

func (d *Deck) GetCards() []ICard {
	return d.Cards
}

func (d *Deck) SetCards(cards []ICard) {
	d.Cards = cards
}
