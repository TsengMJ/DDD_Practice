package deck

type ICard interface {
	Compare() int
	GetType() string
	GetValue() string
}

type Card struct {
	ICard
	Type  string
	Value string
}

func (c *Card) GetType() string {
	return c.Type
}

func (c *Card) GetValue() string {
	return c.Value
}

func (c *Card) Compare(card ICard) int {
	return 0
}
