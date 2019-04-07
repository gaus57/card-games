package pharaoh

type Player struct {
	id      int
	game    *Game
	hand    *Hand
	isTaken bool
}

func newPlayer(id int, g *Game) *Player {
	return &Player{
		id:      id,
		game:    g,
		hand:    newHand(),
		isTaken: false,
	}
}

func (p *Player) canMove() bool {
	if p.game.currentPlayerId != p.id || p.game.isCompleted {
		return false
	}
	for _, card := range p.hand.cards {
		if p.game.checkCardMove(card) {
			return true
		}
	}
	return false
}

func (p *Player) canTake() bool {
	if p.game.currentPlayerId != p.id || p.game.isCompleted {
		return false
	}
	return !p.isTaken && !p.canMove()
}

func (p *Player) canSkip() bool {
	if p.game.currentPlayerId != p.id || p.game.isCompleted {
		return false
	}
	return !p.canTake() && !p.canMove()
}
