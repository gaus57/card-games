package main

import (
	"card-games/decks/classic36"
	"card-games/games/pharaoh"
)

func main() {
	var dec = classic36.Deck()
	var cards = make([]pharaoh.Card, len(dec))
	for i, v := range dec {
		cards[i] = pharaoh.Card(v)
	}
	game := pharaoh.NewGame()
	game.Join()
	game.Join()
	game.Start(cards)
}
