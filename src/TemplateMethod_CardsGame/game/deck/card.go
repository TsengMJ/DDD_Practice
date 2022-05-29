package deck

type ICard interface {
	Compare(card ICard) int
	SetType(cardType string)
	SetValue(cardValue string)
	GetType() string
	GetValue() string
}

type Card struct {
	ICard
	Type  string
	Value string
}

func (c *Card) Compare(card ICard) int {
	return 0
}

func (c *Card) SetType(cardType string) {
	c.Type = cardType
}

func (c *Card) SetValue(cardValue string) {
	c.Value = cardValue
}

func (c *Card) GetType() string {
	return c.Type
}

func (c *Card) GetValue() string {
	return c.Value
}
