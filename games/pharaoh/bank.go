package pharaoh

import (
	"math/rand"
	"time"
)

type Bank struct {
	cards []*GameCard
}

func newBank() *Bank {
	return &Bank{
		cards: make([]*GameCard, 0),
	}
}

func (b *Bank) push(gc *GameCard) {
	b.cards = append(b.cards, gc)
}

func (b *Bank) take() *GameCard {
	if len(b.cards) > 0 {
		gc := b.cards[len(b.cards)-1]
		b.cards = b.cards[:len(b.cards)-1]
		return gc
	}
	return nil
}

func (b *Bank) count() int {
	return len(b.cards)
}

func (b *Bank) mix() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(b.cards), func(i, j int) {
		b.cards[i], b.cards[j] = b.cards[j], b.cards[i]
	})
}
