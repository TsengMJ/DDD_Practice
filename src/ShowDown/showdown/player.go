package showdown

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type IPlayer interface {
	GetPoint() int
	GetOrder() int
	GetDrawCard() DrawCard
	SetName()
	SetOrder(order int)
	SetDrawCard(drawCard DrawCard)
	AddPoint()
	Draw(deck Deck)
	Show() Card
	Exchange() int
}

type Player struct {
	Name     string
	Order    int
	Point    int
	DrawCard DrawCard
}

type HumanPlayer struct {
	Player
}

type AIPlayer struct {
	Player
}

func (player *Player) GetPoint() int {
	return player.Point
}

func (player *Player) GetOrder() int {
	return player.Order
}

func (player *Player) GetDrawCard() DrawCard {
	return player.DrawCard
}

func (player *Player) SetName() {
	fmt.Printf("=================== Player %d ===================\n", player.Order)
	fmt.Printf("Please enter your name: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading player name.", err)
	}

	name := strings.TrimSpace(input)

	player.Name = name
}

func (player *Player) SetOrder(order int) {
	player.Order = order
}

func (player *Player) SetDrawCard(drawCard DrawCard) {
	player.DrawCard = drawCard
}

func (player *Player) AddPoint() {
	player.Point += 1
}

func (player *Player) Draw(deck Deck) {
	newDrawCard := DrawCard{}
	newDrawCard.SetHand(deck, player.Order)
	player.DrawCard = newDrawCard
}

func (player *HumanPlayer) Show() Card {
	fmt.Printf("\n=================== Player %d ===================\n\n", player.Order)
	fmt.Println("Your Cards:")
	for index, card := range player.DrawCard.Hand {
		fmt.Printf("  %d: %s\n", index, card)
	}

	var showCard Card

	for true {
		fmt.Printf("\nPlease enter a card index to show: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
		}

		inputIndex, err := strconv.ParseInt(input, 10, 64)
		targetIndex := int(inputIndex)

		if err != nil {
			fmt.Println("An error occured while parse input to int", err)
		}

		if targetIndex >= 0 && targetIndex < len(player.DrawCard.Hand) {
			showCard = player.DrawCard.Hand[targetIndex]
			player.DrawCard.RemoveCard(targetIndex)
			break
		}

		fmt.Println("Invalid Index, please enter again.")
	}

	fmt.Println("Now Your Cards:")
	for index, card := range player.DrawCard.Hand {
		fmt.Printf("  %d: %s\n", index, card)
	}

	return showCard
}

func (player *AIPlayer) Show() Card {
	cardIndex := rand.Intn(len(player.DrawCard.Hand))
	showCard := player.DrawCard.Hand[cardIndex]

	player.DrawCard.RemoveCard(cardIndex)

	return showCard
}

func (player *HumanPlayer) Exchange() int {

	var playerIndex int = -1

	for true {
		fmt.Println("Enter which Player you want to exchange: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		inputInt, _ := strconv.ParseInt(input, 10, 64)
		playerIndex = int(inputInt)

		if playerIndex >= -1 && playerIndex < 4 && playerIndex != player.Order {
			break
		}
	}

	return playerIndex
}

func (player *AIPlayer) Exchange() int {
	var toPlayerIndex int

	for true {
		randomIndex := rand.Intn(5) - 1

		if randomIndex != player.Order {
			toPlayerIndex = randomIndex
			break
		}
	}

	return toPlayerIndex
}
