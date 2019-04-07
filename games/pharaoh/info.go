package pharaoh

type Info struct {
	players         []PlayerInfo
	currentPlayerId int
	bancCount       int
	discardCount    int
	isCompleted     bool
	hand            []*GameCard
	canMove         bool
	canTake         bool
	canSkip         bool
}

type PlayerInfo struct {
	id        int
	handCount int
	points    int
}

func newInfo(g *Game, playerId int) *Info {
	player := g.getPlayer(playerId)
	if player == nil {
		return nil
	}
	var players []PlayerInfo
	for i, p := range g.players {
		players = append(players, PlayerInfo{
			id:        p.id,
			handCount: p.hand.count(),
			points:    0,
		})
		if g.isCompleted {
			players[i].points = p.hand.points()
		}
	}
	var hand []*GameCard
	for _, card := range player.hand.cards {
		hand = append(hand, card)
	}
	return &Info{
		players:         players,
		currentPlayerId: g.currentPlayerId,
		bancCount:       g.bank.count(),
		discardCount:    g.discard.count(),
		isCompleted:     g.isCompleted,
		hand:            hand,
		canMove:         player.canMove(),
		canTake:         player.canTake(),
		canSkip:         player.canSkip(),
	}
}
