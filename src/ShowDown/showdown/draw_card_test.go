package showdown

import "testing"

func TestSetHand(t *testing.T) {
	deck := Deck{}
	deck.Init()

	for order := 0; order < 4; order++ {
		drawCard := DrawCard{}
		drawCard.SetHand(deck, order)

		for index, card := range drawCard.Hand {
			cardIndex := index + order*13
			if card.rank != deck.cards[cardIndex].rank || card.suit != deck.cards[cardIndex].suit {
				t.Errorf("DrawCard: user get wrong cards")
			}
		}
	}
}

func TestRemoveCard(t *testing.T) {
	deck := Deck{}
	deck.Init()
	drawCard := DrawCard{}
	drawCard.SetHand(deck, 0)

	index := 5
	targetCard := drawCard.Hand[index]
	drawCard.RemoveCard(index)

	if len(drawCard.Hand) > 12 {
		t.Errorf("DrawCard: failed to remove the card")
	}

	if len(drawCard.Hand) < 12 {
		t.Errorf("DrawCard: remove more than one cards")
	}

	for _, card := range drawCard.Hand {
		if card.rank == targetCard.rank && card.suit == targetCard.suit {
			t.Errorf("DrawCard: remove another wrong card")
		}
	}

}
