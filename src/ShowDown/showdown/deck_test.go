package showdown

import (
	"sort"
	"testing"
)

func TestInit(t *testing.T) {
	deck := Deck{}
	deck.Init()

	expectNumberOfCards := 52
	if len(deck.cards) != expectNumberOfCards {
		t.Errorf("Deck: Number of cards %d, wanted %d", len(deck.cards), expectNumberOfCards)
	}

	rankKeys := make([]string, 0)
	suitKeys := make([]string, 0)

	for rankKey, _ := range RankTable {
		rankKeys = append(rankKeys, rankKey)
	}

	for suitKey, _ := range SuitTable {
		rankKeys = append(suitKeys, suitKey)
	}

	sort.Strings(rankKeys)
	sort.Strings(suitKeys)

	for i, rankKey := range rankKeys {
		for j, suitKey := range suitKeys {
			index := i + j*13

			if deck.cards[index].rank != rankKey || deck.cards[index].suit != suitKey {
				t.Errorf("Deck: Does not create the cards corretly")
				return
			}

		}
	}
}
