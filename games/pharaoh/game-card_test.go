package pharaoh

import "testing"

type MockCard struct {
	cod  string
	suit string
}

func newTestGameCard(cod string, suit string) *GameCard {
	return &GameCard{
		card: MockCard{cod: cod, suit: suit},
	}
}

func (c MockCard) Code() string {
	return c.cod
}

func (c MockCard) SuitCode() string {
	return c.suit
}

func TestCode(t *testing.T) {
	gc := newTestGameCard("10", "diamonds")
	if gc.code() != "10-diamonds" {
		t.Error("Error game card code")
	}
}

func TestPoints(t *testing.T) {
	gc := newTestGameCard("10", "diamonds")
	if gc.points() != 10 {
		t.Error("Error game card points")
	}
	gc = newTestGameCard("queen", "diamonds")
	if gc.points() != 30 {
		t.Error("Error game card points")
	}
	gc = newTestGameCard("invalid", "diamonds")
	if gc.points() != 0 {
		t.Error("Error invalid game card points")
	}
}
