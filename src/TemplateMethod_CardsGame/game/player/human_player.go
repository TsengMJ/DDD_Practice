package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"template_method/game/deck"
)

type HumanPlayer struct {
	Player
}

func (p *AIPlayer) HumanPlayer() deck.ICard {

	var showCard deck.ICard
	for true {
		fmt.Printf("\nPlease enter a card index to show: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
		}

		inIdx, err := strconv.ParseInt(input, 10, 64)
		idx := int(inIdx)

		if err != nil {
			fmt.Println("An error occured while parse input to int", err)
		}

		if idx >= 0 && idx < p.GetNumberOfCards() {
			showCard = p.Cards[idx]
			p.removeCard(idx)
			break
		}

		fmt.Println("Invalid Index, please enter again.")
	}

	return showCard
}
