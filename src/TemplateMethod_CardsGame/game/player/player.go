package player

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"template_method/game/deck"
)

type IPlayer interface {
	Rename()
	DrawCard(deck *deck.Deck)
	SetOrder(order int)
	ShowCard() deck.Card
	AddPoint()
	RemoveCard(index int)
	NumberOfHands() int
	GetPoint() int
}

type Player struct {
	IPlayer
	Name  string
	Point int
	Order int
	Hands []deck.Card
}

func (p *Player) Rename() {
	fmt.Printf("=================== Player %d ===================\n", p.Order)
	fmt.Printf("Please enter your name: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading player name.", err)
	}

	name := strings.TrimSpace(input)

	p.Name = name
}

func (p *Player) SetOrder(order int) {
	p.Order = order
}

func (p *Player) DrawCard(deck *deck.Deck) {
	p.Hands = append(p.Hands, deck.Draw())
}

func (p *Player) AddPoint() {
	p.Point += 1
}

func (p *Player) RemoveCard(index int) {
	p.Hands = append(p.Hands[:index], p.Hands[index+1:]...)
}

func (p *Player) NumberOfHands() int {
	return len(p.Hands)
}

func (p *Player) GetPoint() int {
	return p.Point
}
