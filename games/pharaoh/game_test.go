package pharaoh

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame()
	if len(g.players) != 0 {
		t.Error("Error new game")
	}
	if g.isCompleted {
		t.Error("Error new game")
	}
	if g.currentPlayerId != 0 {
		t.Error("Error new game")
	}
}

func TestGameJoin(t *testing.T) {
	g := NewGame()
	id, err := g.Join()
	if id != 0 || err != nil || len(g.players) != 1 {
		t.Error("Error join game")
	}
	id, err = g.Join()
	if id != 1 || err != nil || len(g.players) != 2 {
		t.Error("Error join game")
	}
	g.Join()
	g.Join()
	g.Join()
	id, err = g.Join()
	if id != -1 || err == nil {
		t.Error("Error join game")
	}
}

func TestGameDealCard(t *testing.T) {
	g := NewGame()
	g.Join()
	p := g.players[0]
	g.dealCard(p, 1)
	if g.bank.count() != 0 || p.hand.count() != 0 {
		t.Error("Error game deal card")
	}
	g.bank.push(newTestGameCard("10", "clubs"))
	g.bank.push(newTestGameCard("queen", "clubs"))
	g.dealCard(p, 1)
	if g.bank.count() != 1 || p.hand.count() != 1 {
		t.Error("Error game deal card")
	}
	g.dealCard(p, 2)
	if g.bank.count() != 0 || p.hand.count() != 2 {
		t.Error("Error game deal card")
	}
	g.discard.push(newTestGameCard("9", "clubs"))
	g.discard.push(newTestGameCard("8", "clubs"))
	g.discard.push(newTestGameCard("7", "clubs"))
	g.discard.push(newTestGameCard("6", "clubs"))
	g.dealCard(p, 1)
	if g.bank.count() != 2 || p.hand.count() != 3 || g.discard.count() != 1 {
		t.Error("Error game deal card")
	}
	if g.discard.cards[0].code() != "6-clubs" {
		t.Error("Error game deal card")
	}
}

func TestGameStart(t *testing.T) {
	deck := []Card{
		MockCard{cod: "6", suit: "clubs"},
		MockCard{cod: "7", suit: "clubs"},
		MockCard{cod: "8", suit: "clubs"},
		MockCard{cod: "9", suit: "clubs"},
		MockCard{cod: "10", suit: "clubs"},
		MockCard{cod: "6", suit: "hearts"},
		MockCard{cod: "7", suit: "hearts"},
		MockCard{cod: "8", suit: "hearts"},
		MockCard{cod: "9", suit: "hearts"},
		MockCard{cod: "10", suit: "hearts"},
	}
	g := NewGame()
	if g.Start(deck) == nil {
		t.Error("Error start game")
	}
	g.Join()
	if g.Start(deck) == nil {
		t.Error("Error start game")
	}
	g.Join()
	if g.Start(deck) != nil {
		t.Error("Error start game")
	}
	if !g.isStarted {
		t.Error("Error start game")
	}
	if g.players[0].hand.count() != 4 || g.players[1].hand.count() != 4 {
		t.Error("Error start game")
	}
	if g.bank.count() != 2 {
		t.Error("Error start game")
	}
}

func TestGameGetPlayer(t *testing.T) {
	g := NewGame()
	if g.getPlayer(0) != nil {
		t.Error("Error get player")
	}
	g.Join()
	g.Join()
	if g.getPlayer(3) != nil {
		t.Error("Error get player")
	}
	if g.getPlayer(1) == nil || g.getPlayer(1).id != 1 {
		t.Error("Error get player")
	}
}

func TestGameNextPlayer(t *testing.T) {
	g := NewGame()
	if g.nextPlayer() != nil {
		t.Error("Error game next player")
	}
	g.Join()
	g.Join()
	g.Join()
	next := g.nextPlayer()
	if next == nil || next.id != 1 || g.currentPlayerId != 1 {
		t.Error("Error game next player")
	}
	next.isTaken = true
	next = g.nextPlayer()
	if next == nil || next.id != 2 || g.currentPlayerId != 2 {
		t.Error("Error game next player")
	}
	if g.players[1].isTaken {
		t.Error("Error game next player")
	}
	next = g.nextPlayer()
	if next == nil || next.id != 0 || g.currentPlayerId != 0 {
		t.Error("Error game next player")
	}
}

func TestGameTopCard(t *testing.T) {
	g := NewGame()
	if g.topCard() != nil {
		t.Error("Error game top card")
	}
	g.discard.push(newTestGameCard("10", "clubs"))
	if g.topCard() == nil || g.topCard().code() != "10-clubs" {
		t.Error("Error game top card")
	}
	g.discard.push(newTestGameCard("6", "clubs"))
	if g.topCard() == nil || g.topCard().code() != "6-clubs" {
		t.Error("Error game top card")
	}
	if g.discard.count() != 2 {
		t.Error("Error game top card")
	}
}

func TestGameCheckComplete(t *testing.T) {
	g := NewGame()
	g.checkComplete()
	if g.isCompleted {
		t.Error("Error game check complete")
	}
	g.Join()
	g.checkComplete()
	if g.isCompleted {
		t.Error("Error game check complete")
	}
	g.Join()
	g.isStarted = true
	g.players[0].hand.push(newTestGameCard("10", "clubs"))
	g.players[1].hand.push(newTestGameCard("9", "clubs"))
	g.checkComplete()
	if g.isCompleted {
		t.Error("Error game check complete")
	}
	g.players[0].hand.give("10-clubs")
	g.checkComplete()
	if !g.isCompleted {
		t.Error("Error game check complete")
	}
}

func TestGameCheckCardMove(t *testing.T) {
	g := NewGame()
	if !g.checkCardMove(newTestGameCard("10", "clubs")) {
		t.Error("Error game check card move")
	}
	g.discard.push(newTestGameCard("6", "diamonds"))
	if g.checkCardMove(newTestGameCard("10", "clubs")) {
		t.Error("Error game check card move")
	}
	if !g.checkCardMove(newTestGameCard("6", "clubs")) {
		t.Error("Error game check card move")
	}
	if !g.checkCardMove(newTestGameCard("10", "diamonds")) {
		t.Error("Error game check card move")
	}
	if !g.checkCardMove(newTestGameCard("queen", "hearts")) {
		t.Error("Error game check card move")
	}
	g.discard.push(newTestGameCard("queen", "diamonds"))
	g.requestedSuitCode = "hearts"
	if g.checkCardMove(newTestGameCard("10", "diamonds")) {
		t.Error("Error game check card move")
	}
	if !g.checkCardMove(newTestGameCard("10", "hearts")) {
		t.Error("Error game check card move")
	}
}

func TestGamePutCard(t *testing.T) {
	deck := []Card{
		MockCard{cod: "6", suit: "clubs"},
		MockCard{cod: "7", suit: "clubs"},
		MockCard{cod: "8", suit: "clubs"},
		MockCard{cod: "9", suit: "clubs"},
		MockCard{cod: "10", suit: "clubs"},
		MockCard{cod: "6", suit: "hearts"},
		MockCard{cod: "7", suit: "hearts"},
		MockCard{cod: "8", suit: "hearts"},
		MockCard{cod: "9", suit: "hearts"},
		MockCard{cod: "10", suit: "hearts"},
		MockCard{cod: "queen", suit: "hearts"},
		MockCard{cod: "king", suit: "hearts"},
	}
	g := NewGame()
	g.Join()
	g.Join()
	if g.Start(deck) != nil {
		t.Error("Error game put card")
	}
	if g.putCard(newTestGameCard("10", "hearts"), "clubs") != nil {
		t.Error("Error game put card")
	}
	if g.requestedSuitCode != "" {
		t.Error("Error game put card")
	}
	if g.currentPlayerId != 1 {
		t.Error("Error game put card")
	}
	if g.putCard(newTestGameCard("9", "clubs"), "") == nil {
		t.Error("Error game put card")
	}
	if g.putCard(newTestGameCard("7", "hearts"), "") != nil {
		t.Error("Error game put card")
	}
	if g.currentPlayerId != 1 {
		t.Error("Error game put card 7")
	}
	if g.players[0].hand.count() != 6 || g.bank.count() != 2 {
		t.Error("Error game put card 7")
	}
	if g.putCard(newTestGameCard("queen", "clubs"), "diamonds") != nil {
		t.Error("Error game put card queen")
	}
	if g.requestedSuitCode != "diamonds" {
		t.Error("Error game put card queen")
	}
	if g.putCard(newTestGameCard("8", "diamonds"), "") != nil {
		t.Error("Error game put card")
	}
}

func TestGameMove(t *testing.T) {
	g := NewGame()
	g.bank.push(newTestGameCard("queen", "clubs"))
	g.bank.push(newTestGameCard("6", "hearts"))
	g.bank.push(newTestGameCard("10", "hearts"))
	g.bank.push(newTestGameCard("9", "hearts"))
	g.bank.push(newTestGameCard("7", "spades"))
	g.Join()
	g.players[0].hand.push(newTestGameCard("10", "diamonds"))
	g.players[0].hand.push(newTestGameCard("9", "clubs"))
	g.Join()
	g.players[1].hand.push(newTestGameCard("8", "hearts"))
	g.players[1].hand.push(newTestGameCard("queen", "hearts"))
	if g.Move(&Move{playerId: 0, cardCode: "9-clubs"}) == nil {
		t.Error("Error game move")
	}
	g.isStarted = true
	if g.Move(&Move{playerId: 0, cardCode: "9-hearts"}) == nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 0, cardCode: "9-clubs"}) != nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 2, takeCard: true}) == nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 0, takeCard: true}) == nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 1, takeCard: true}) == nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 1, cardCode: "8-hearts"}) == nil {
		t.Error("Error game move")
	}
	if g.players[1].hand.count() != 2 {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 1, cardCode: "queen-hearts", suitCode: "spades"}) != nil {
		t.Error("Error game move")
	}
	if g.requestedSuitCode != "spades" {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 0, takeCard: true}) != nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 0, cardCode: "7-spades"}) != nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 0, takeCard: true}) != nil {
		t.Error("Error game move")
	}
	if g.currentPlayerId != 1 {
		t.Error("Error game move")
	}
	if g.Move(&Move{playerId: 1}) == nil {
		t.Error("Error game move")
	}
}

func TestGameInfo(t *testing.T) {
	g := NewGame()
	if g.Info(1) != nil {
		t.Error("Error game info")
	}
	g.bank.push(newTestGameCard("queen", "clubs"))
	g.bank.push(newTestGameCard("6", "hearts"))
	g.bank.push(newTestGameCard("10", "hearts"))
	g.Join()
	g.players[0].hand.push(newTestGameCard("10", "diamonds"))
	g.players[0].hand.push(newTestGameCard("9", "clubs"))
	g.Join()
	g.players[1].hand.push(newTestGameCard("8", "hearts"))
	g.isStarted = true
	g.currentPlayerId = 1
	info := g.Info(1)
	if info == nil {
		t.Error("Error game info")
	}
	if info.players[0].handCount != 2 || info.players[1].handCount != 1 {
		t.Error("Error game info")
	}
	if info.players[0].points != 0 || info.players[1].points != 0 {
		t.Error("Error game info")
	}
	if info.currentPlayerId != 1 {
		t.Error("Error game info")
	}
	if info.bancCount != 3 {
		t.Error("Error game info")
	}
	if info.discardCount != 0 {
		t.Error("Error game info")
	}
	if info.hand[0].code() != "8-hearts" {
		t.Error("Error game info")
	}
	if info.isCompleted {
		t.Error("Error game info")
	}
	if !info.canMove {
		t.Error("Error game info")
	}
	if info.canTake {
		t.Error("Error game info")
	}
	if info.canSkip {
		t.Error("Error game info")
	}
	g.isCompleted = true
	info = g.Info(0)
	if info == nil {
		t.Error("Error game info")
	}
	if !info.isCompleted {
		t.Error("Error game info")
	}
	if info.players[0].points != 19 || info.players[1].points != 8 {
		t.Error("Error game info")
	}
}
