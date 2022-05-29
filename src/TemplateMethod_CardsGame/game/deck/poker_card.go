package deck

type PokerCard struct {
	Card
}

var PokerType = map[string]int{
	"Spade":   0,
	"Heart":   1,
	"Diamond": 2,
	"Club":    3,
}

var PokerValue = map[string]int{
	"2":  0,
	"3":  1,
	"4":  2,
	"5":  3,
	"6":  4,
	"7":  5,
	"8":  6,
	"9":  7,
	"10": 8,
	"J":  9,
	"Q":  10,
	"K":  11,
	"A":  12,
}

func (c *PokerCard) Compare(card PokerCard) int {
	if PokerType[c.Type] > PokerType[card.Type] || (PokerType[c.Type] == PokerType[card.Type] && PokerValue[c.Value] > PokerValue[card.Value]) {
		return 1
	} else {
		return -1
	}
}
