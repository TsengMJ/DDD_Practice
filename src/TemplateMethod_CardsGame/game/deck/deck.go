package deck

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
	Next  int
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) Draw() Card {
	card := d.Cards[d.Next]
	d.Next += 1
	return card
}
