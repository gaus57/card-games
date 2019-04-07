package pharaoh

import "testing"

func TestNewPlayer(t *testing.T) {
	g := NewGame()
	p := newPlayer(2, g)
	if p.id != 2 {
		t.Error("Error new player")
	}
	if p.game != g {
		t.Error("Error new player")
	}
	if p.hand.count() != 0 {
		t.Error("Error new player")
	}
	if p.isTaken {
		t.Error("Error new player")
	}
}

func TestPlayerCanMove(t *testing.T) {
	g := NewGame()
	p := newPlayer(0, g)
	if p.canMove() {
		t.Error("Error player can move")
	}
	g.players = append(g.players, p)
	g.currentPlayerId = 0
	p.hand.push(newTestGameCard("10", "clubs"))
	if !p.canMove() {
		t.Error("Error player can move")
	}
	g.isCompleted = true
	if p.canMove() {
		t.Error("Error player can move")
	}
}

func TestPlayerCanTake(t *testing.T) {
	g := NewGame()
	g.currentPlayerId = 1
	p := newPlayer(0, g)
	if p.canTake() {
		t.Error("Error player can take")
	}
	g.players = append(g.players, p)
	g.currentPlayerId = 0
	if !p.canTake() {
		t.Error("Error player can take")
	}
	g.isCompleted = true
	if p.canTake() {
		t.Error("Error player can take")
	}
	g.isCompleted = false
	p.isTaken = true
	if p.canTake() {
		t.Error("Error player can take")
	}
	p.isTaken = false
	p.hand.push(newTestGameCard("10", "clubs"))
	if p.canTake() {
		t.Error("Error player can take")
	}
}

func TestPlayerCanSkip(t *testing.T) {
	g := NewGame()
	g.currentPlayerId = 1
	p := newPlayer(0, g)
	if p.canSkip() {
		t.Error("Error player can skip")
	}
	g.players = append(g.players, p)
	g.currentPlayerId = 0
	p.isTaken = true
	if !p.canSkip() {
		t.Error("Error player can skip")
	}
	g.isCompleted = true
	if p.canSkip() {
		t.Error("Error player can skip")
	}
}
