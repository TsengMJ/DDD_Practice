package deck

type UNOCard struct {
	Card
}

type UNOType int

const (
	Blue UNOType = iota
	Red
	Yellow
	Green
)

type UNOValue int

const (
	One UNOValue = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

func (c *UNOCard) Compare(card UNOCard) int {
	if c.Type > card.Type || (c.Type == card.Type && c.Value > card.Value) {
		return 1
	} else {
		return -1
	}
}

func (c *UNOCard) GetAllType() []string {
	return getAllUNOTypes()
}

func (c *UNOCard) GetAllValues() []string {
	return getAllUNOValues()
}

func getAllUNOTypes() []string {
	return []string{"Blue", "Red", "Yellow", "Green"}
}

func getAllUNOValues() []string {
	return []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
}
