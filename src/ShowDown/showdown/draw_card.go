package showdown

type DrawCard struct {
	Hand []Card
}

func (drawCard *DrawCard) SetHand(deck Deck, order int) {
	startIndex := order * 13
	endIndex := startIndex + 13
	drawCard.Hand = deck.cards[startIndex:endIndex]
}

func (drawCard *DrawCard) RemoveCard(index int) {
	drawCard.Hand = append(drawCard.Hand[:index], drawCard.Hand[index+1:]...)
}
