package player

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"template_method/game/deck"
)

type IPlayer interface {
	DrawCards(deck deck.IDeck)
	ShowCard() deck.ICard
	ReName()
	AddPoint()
	AddCard(card deck.ICard)
	GetPoint() int
	GetName() string
	GetNumberOfCards() int
	removeCard(index int)
}

type Player struct {
	Name  string
	Cards []deck.ICard
	Point int
}

func (p *Player) DrawCards(d deck.IDeck) {
	p.Cards = append(p.Cards, d.Draw())
}

func (p *Player) ShowCard() deck.ICard {
	return p.Cards[0]
}

func (p *Player) ReName() {
	fmt.Printf("Please enter player name: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading player name.", err)
	}

	name := strings.TrimSpace(input)

	p.Name = name
}

func (p *Player) AddPoint() {
	p.Point += 1
}

func (p *Player) AddCard(card deck.ICard) {
	p.Cards = append(p.Cards, card)
}

func (p *Player) GetPoint() int {
	return p.Point
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) GetNumberOfCards() int {
	return len(p.Cards)
}

func (p *Player) removeCard(index int) {
	p.Cards = append(p.Cards[:index], p.Cards[index+1:]...)
}
