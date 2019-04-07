package pharaoh

type Hand struct {
	cards map[string]*GameCard
}

func newHand() *Hand {
	return &Hand{
		cards: make(map[string]*GameCard),
	}
}

func (h *Hand) push(gc *GameCard) {
	h.cards[gc.code()] = gc
}

func (h *Hand) give(code string) *GameCard {
	if gc, ok := h.cards[code]; ok {
		delete(h.cards, code)
		return gc
	}
	return nil
}

func (h *Hand) count() int {
	return len(h.cards)
}

func (h *Hand) points() int {
	var points = 0
	for _, gc := range h.cards {
		points += gc.points()
	}
	return points
}
