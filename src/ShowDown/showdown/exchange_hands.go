package showdown

type ExchangeHand struct {
	IsUsed     bool
	FromPlayer Player
	ToPlayer   Player
	Round      int
}

func (exchangeHand *ExchangeHand) ExchangeHands(fromPlayer, toPlayer Player) {
	exchangeHand.FromPlayer = fromPlayer
	exchangeHand.ToPlayer = toPlayer

	exchangeHand.FromPlayer.DrawCard, exchangeHand.ToPlayer.DrawCard = exchangeHand.ToPlayer.DrawCard, exchangeHand.FromPlayer.DrawCard
	exchangeHand.Round = 3
}

func (exchangeHand *ExchangeHand) MinusRound() {
	exchangeHand.Round -= 1

	if exchangeHand.Round == 0 {
		exchangeHand.FromPlayer.DrawCard, exchangeHand.ToPlayer.DrawCard = exchangeHand.ToPlayer.DrawCard, exchangeHand.FromPlayer.DrawCard
	}
}
