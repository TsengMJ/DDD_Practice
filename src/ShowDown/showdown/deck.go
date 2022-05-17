package showdown

import (
	"math/rand"
	"sort"
	"time"
)

type Deck struct {
	cards []Card
}

func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.cards), func(i, j int) { deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i] })
}

func (deck *Deck) Init() {
	rankKeys := make([]string, 0)
	suitKeys := make([]string, 0)

	for rankKey, _ := range RankTable {
		rankKeys = append(rankKeys, rankKey)
	}

	for suitKey, _ := range SuitTable {
		suitKeys = append(suitKeys, suitKey)
	}

	sort.Strings(rankKeys)
	sort.Strings(suitKeys)

	for _, rankKey := range rankKeys {
		for _, suitKey := range suitKeys {
			var newCard = Card{rank: rankKey, suit: suitKey}
			deck.cards = append(deck.cards, newCard)
		}
	}
}
