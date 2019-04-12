package pharaoh

type Game struct {
	MinPlayers        int
	MaxPlayers        int
	players           []*Player
	bank              *Bank
	discard           *Bank
	currentPlayerId   int
	requestedSuitCode string
	isCompleted       bool
	isStarted         bool
}

func NewGame() *Game {
	return &Game{
		MinPlayers:        2,
		MaxPlayers:        2,
		players:           make([]*Player, 0, 2),
		bank:              newBank(),
		discard:           newBank(),
		currentPlayerId:   0,
		requestedSuitCode: "",
		isCompleted:       false,
		isStarted:         false,
	}
}

func (g *Game) Join() (int, error) {
	if len(g.players) < g.MaxPlayers {
		var playerId = len(g.players)
		g.players = append(g.players, newPlayer(playerId, g))
		return playerId, nil
	}
	return -1, newError("Превышено максимальное количество игроков")
}

func (g *Game) dealCard(player *Player, count int) {
	for i := 0; i < count; i++ {
		if card := g.bank.take(); card != nil {
			player.hand.push(card)
		} else if g.discard.count() > 0 {
			g.discard, g.bank = g.bank, g.discard
			g.discard.push(g.bank.take())
			g.bank.mix()
			i--
		} else {
			break
		}
	}
}

func (g *Game) Start(deck []Card) error {
	if len(g.players) < g.MinPlayers {
		return newError("Нехватает игроков")
	}
	g.isStarted = true
	for _, card := range deck {
		g.bank.push(&GameCard{card: card})
	}
	g.bank.mix()
	for _, player := range g.players {
		g.dealCard(player, 4)
	}
	return nil
}

func (g *Game) getPlayer(id int) *Player {
	if id < 0 || id >= len(g.players) {
		return nil
	}
	return g.players[id]
}

func (g *Game) currentPlayer() *Player {
	if len(g.players) > 0 {
		return g.players[g.currentPlayerId]
	}
	return nil
}

func (g *Game) nextPlayer() *Player {
	if g.currentPlayer() == nil {
		return nil
	}
	g.currentPlayer().isTaken = false
	g.currentPlayerId++
	if g.currentPlayerId >= len(g.players) {
		g.currentPlayerId = 0
	}
	return g.currentPlayer()
}

func (g *Game) topCard() *GameCard {
	if g.discard.count() > 0 {
		return g.discard.cards[len(g.discard.cards)-1]
	}
	return nil
}

func (g *Game) checkComplete() {
	if !g.isStarted {
		return
	}
	for _, p := range g.players {
		if p.hand.count() == 0 {
			g.isCompleted = true
			break
		}
	}
}

func (g *Game) checkCardMove(gc *GameCard) bool {
	if gc.card.Code() == "queen" {
		return true
	}
	if g.requestedSuitCode != "" {
		if gc.card.SuitCode() == g.requestedSuitCode {
			return true
		}
		return false
	}
	topCard := g.topCard()
	if topCard == nil {
		return true
	}
	if topCard.card.Code() == gc.card.Code() || topCard.card.SuitCode() == gc.card.SuitCode() {
		return true
	}
	return false
}

func (g *Game) putCard(gc *GameCard, requestedSuitCode string) error {
	if !g.checkCardMove(gc) {
		return newError("Невозможно сходить этой картой")
	}
	g.discard.push(gc)
	g.requestedSuitCode = ""
	var nextPlayer = g.nextPlayer()
	if gc.card.Code() == "7" {
		g.dealCard(nextPlayer, 2)
		g.nextPlayer()
	} else if gc.card.Code() == "queen" {
		g.requestedSuitCode = requestedSuitCode
	}
	return nil
}

func (g *Game) Move(m *Move) error {
	if !g.isStarted || g.isCompleted {
		return newError("Игра неактивна")
	}
	player := g.getPlayer(m.PlayerId)
	if player == nil {
		return newError("Игрок не в игре")
	}
	if g.currentPlayerId != m.PlayerId {
		return newError("Невозможно играть не в свой ход")
	}
	if m.CardCode != "" {
		card := player.hand.give(m.CardCode)
		if card == nil {
			return newError("У игрока нет такой карты")
		}
		if err := g.putCard(card, m.SuitCode); err != nil {
			player.hand.push(card)
			return err
		}
	} else if m.TakeCard {
		if !player.canTake() {
			return newError("Невозможно взять карту")
		}
		g.dealCard(player, 1)
		player.isTaken = true
		if !player.canMove() {
			g.nextPlayer()
		}
	} else {
		return newError("Не выбрано ниодного действия")
	}
	g.checkComplete()
	return nil
}

func (g *Game) Info(playerId int) *Info {
	return newInfo(g, playerId)
}
