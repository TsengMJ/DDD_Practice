package showdown

import "fmt"

type Showdown struct {
	Players       []IPlayer
	Deck          Deck
	Round         int
	ExchangeHands []ExchangeHand
}

func (showdown *Showdown) StartGame(players []IPlayer, deck Deck) {
	fmt.Println("=================== Game start ===================\n")

	showdown.Players = players
	showdown.Deck = deck
	showdown.Round = 0
	showdown.Deck.Shuffle()

	for index, player := range showdown.Players {
		player.SetOrder(index)
		player.SetName()
		player.Draw(deck)
	}

	for i := 0; i < 4; i++ {
		showdown.ExchangeHands = append(showdown.ExchangeHands, ExchangeHand{})
	}

	showdown.NextRound()
}

func (showdown *Showdown) NextRound() {
	if showdown.Round < 13 {
		fmt.Printf("\n=================== Round %d ===================\n", showdown.Round)

		// check

		// check if any player want to exchange
		for _, player := range showdown.Players {
			if !showdown.ExchangeHands[player.GetOrder()].IsUsed {
				toPlayerIndex := player.Exchange()

				if toPlayerIndex != -1 {
					fromDrawCard := showdown.Players[player.GetOrder()].GetDrawCard()
					toDrawCARD := showdown.Players[toPlayerIndex].GetDrawCard()

					showdown.Players[player.GetOrder()].SetDrawCard(toDrawCARD)
					showdown.Players[toPlayerIndex].SetDrawCard(fromDrawCard)

					showdown.ExchangeHands[player.GetOrder()].IsUsed = true
					showdown.ExchangeHands[player.GetOrder()].FromPlayer = player
				}
			}
		}

		// get cards from plyaers
		cards := make([]Card, 0)
		for _, player := range showdown.Players {
			cards = append(cards, player.Show())
		}

		// compare cards and add point to player
		for index, card := range cards {
			fmt.Printf("\nPlayer %d show %s\n", index, card)
		}

		winnerIndex := CompareCards(cards)
		showdown.Players[winnerIndex].AddPoint()

		showdown.AddRound()
		showdown.NextRound()
	} else {
		// end game
		fmt.Println("\n=================== Game end ===================\n")
		for _, player := range showdown.Players {
			fmt.Printf("Player %d get %d points\n", player.GetOrder(), player.GetPoint())
		}

	}
}

func (showdown *Showdown) AddRound() {
	showdown.Round += 1
}

func CompareCards(cards []Card) int {
	largeetCardIndex := 0
	tmpCard := cards[0]

	for index, card := range cards {
		if RankTable[card.rank] > RankTable[tmpCard.rank] ||
			(RankTable[card.rank] == RankTable[tmpCard.rank] && SuitTable[card.suit] > SuitTable[tmpCard.suit]) {
			largeetCardIndex = index
			tmpCard = card
		}
	}

	return largeetCardIndex
}
