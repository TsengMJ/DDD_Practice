package deck

type ICard interface {
	Compare(card Card) int
}

type Card struct {
	ICard
	Type  string
	Value string
}

// func (c *Card) Compare(card Card) int {
// 	return 0
// }
