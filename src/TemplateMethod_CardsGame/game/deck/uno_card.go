package deck

type UnoCard struct {
	Card
}

var UnoType = map[string]int{
	"Blue":   0,
	"Red":    1,
	"Yellow": 2,
	"Green":  3,
}

var UnoValue = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func (c *UnoCard) Compare(card ICard) int {
	if UnoType[c.GetType()] == UnoType[card.GetType()] || UnoValue[c.GetValue()] == UnoValue[card.GetValue()] {
		return 1
	} else {
		return -1
	}
}
