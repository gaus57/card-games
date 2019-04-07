package pharaoh

import "testing"

func TestNewBank(t *testing.T) {
	b := newBank()
	if len(b.cards) != 0 {
		t.Error("New bank not empty")
	}
}

func TestBankPush(t *testing.T) {
	b := newBank()
	var gc = newTestGameCard("10", "diamonds")
	b.push(gc)
	if len(b.cards) != 1 {
		t.Error("Error push card to bank")
	}
	if b.cards[0] != gc {
		t.Error("Error push card to bank")
	}
}

func TestBankTake(t *testing.T) {
	b := newBank()
	gc := newTestGameCard("10", "diamonds")
	b.push(gc)
	actual := b.take()
	if len(b.cards) != 0 {
		t.Error("Error take card from bank")
	}
	if actual != gc {
		t.Error("Error take card from bank")
	}
	if b.take() != nil {
		t.Error("Error take card from empty bank")
	}
}

func TestBankCount(t *testing.T) {
	b := newBank()
	if b.count() != 0 {
		t.Error("Error count bank")
	}
	gc := newTestGameCard("10", "diamonds")
	b.push(gc)
	if b.count() != 1 {
		t.Error("Error count bank")
	}
}

func TestBankMix(t *testing.T) {
	b := newBank()
	gc := newTestGameCard("10", "diamonds")
	b.push(gc)
	b.push(gc)
	b.push(gc)
	b.mix()
	if b.count() != 3 {
		t.Error("Error mix bank")
	}
}
