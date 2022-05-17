package showdown

import (
	"testing"
)

func TestSetName(t *testing.T) {
	player := Player{}
	newName := "MJ"
	player.SetName()

	if player.Name != newName {
		t.Errorf("Player: failed to rename himself")
	}
}

func TestSetOrder(t *testing.T) {
	player := Player{}
	newOrder := 3
	player.SetOrder(newOrder)

	if player.Order != newOrder {
		t.Errorf("Player: failed to set order")
	}
}

func TestAddPoint(t *testing.T) {
	player := Player{}
	player.AddPoint()

	if player.Point != 1 {
		t.Errorf("Player: failed to add point")
	}
}

func TestDraw(t *testing.T) {
	deck := Deck{}
	deck.Init()

	player := Player{}
	player.Draw(deck)

	if len(player.DrawCard.Hand) != 13 {
		t.Errorf("Player: failed to draw card")
	}
}
