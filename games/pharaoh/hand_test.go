package pharaoh

import "testing"

func TestNewHand(t *testing.T) {
	h := newHand()
	if len(h.cards) != 0 {
		t.Error("Error new hand")
	}
}

func TestHandPush(t *testing.T) {
	h := newHand()
	gc := newTestGameCard("10", "diamonds")
	h.push(gc)
	if len(h.cards) != 1 {
		t.Error("Error push hand")
	}
	if actual, ok := h.cards["10-diamonds"]; ok {
		if actual != gc {
			t.Error("Error push hand")
		}
	} else {
		t.Error("Error push hand")
	}
}

func TestHandGive(t *testing.T) {
	h := newHand()
	gc := newTestGameCard("10", "diamonds")
	h.push(gc)
	if h.give("invalid") != nil {
		t.Error("Error give invalid card from hand")
	}
	if h.give("10-diamonds") == gc {
		if len(h.cards) != 0 {
			t.Error("Error give card from hand")
		}
	} else {
		t.Error("Error give card from hand")
	}
}

func TestHandCount(t *testing.T) {
	h := newHand()
	h.push(newTestGameCard("10", "diamonds"))
	h.push(newTestGameCard("6", "diamonds"))
	if h.count() != 2 {
		t.Error("Error count hand")
	}
}

func TestHandPoints(t *testing.T) {
	h := newHand()
	if h.points() != 0 {
		t.Error("Error hand points")
	}
	h.push(newTestGameCard("10", "diamonds"))
	h.push(newTestGameCard("queen", "diamonds"))
	if h.points() != 40 {
		t.Error("Error hand points")
	}
}
