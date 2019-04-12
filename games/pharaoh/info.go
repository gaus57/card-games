package pharaoh

type Info struct {
	PlayerId        int
	Players         []PlayerInfo
	CurrentPlayerId int
	BancCount       int
	DiscardCount    int
	IsCompleted     bool
	TopCard         CardInfo
	Hand            []CardInfo
	CanMove         bool
	CanTake         bool
	CanSkip         bool
}

type PlayerInfo struct {
	Id        int
	HandCount int
	Points    int
}

type CardInfo struct {
	Code     string
	CardCode string
	SuitCode string
}

func newInfo(g *Game, playerId int) *Info {
	player := g.getPlayer(playerId)
	if player == nil {
		return nil
	}
	var players []PlayerInfo
	for i, p := range g.players {
		players = append(players, PlayerInfo{
			Id:        p.id,
			HandCount: p.hand.count(),
			Points:    0,
		})
		if g.isCompleted {
			players[i].Points = p.hand.points()
		}
	}
	var hand []CardInfo
	for _, card := range player.hand.cards {
		gc := CardInfo{
			Code:     card.code(),
			CardCode: card.card.Code(),
			SuitCode: card.card.SuitCode(),
		}
		hand = append(hand, gc)
	}
	info := &Info{
		PlayerId:        playerId,
		Players:         players,
		CurrentPlayerId: g.currentPlayerId,
		BancCount:       g.bank.count(),
		DiscardCount:    g.discard.count(),
		IsCompleted:     g.isCompleted,
		Hand:            hand,
		CanMove:         player.canMove(),
		CanTake:         player.canTake(),
		CanSkip:         player.canSkip(),
	}
	tc := g.topCard()
	if tc != nil {
		info.TopCard = CardInfo{
			Code:     tc.code(),
			CardCode: tc.card.Code(),
			SuitCode: tc.card.SuitCode(),
		}
	}
	return info
}
