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
	if g.Move(&Move{PlayerId: 0, CardCode: "9-clubs"}) == nil {
		t.Error("Error game move")
	}
	g.isStarted = true
	if g.Move(&Move{PlayerId: 0, CardCode: "9-hearts"}) == nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 0, CardCode: "9-clubs"}) != nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 2, TakeCard: true}) == nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 0, TakeCard: true}) == nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 1, TakeCard: true}) == nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 1, CardCode: "8-hearts"}) == nil {
		t.Error("Error game move")
	}
	if g.players[1].hand.count() != 2 {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 1, CardCode: "queen-hearts", SuitCode: "spades"}) != nil {
		t.Error("Error game move")
	}
	if g.requestedSuitCode != "spades" {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 0, TakeCard: true}) != nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 0, CardCode: "7-spades"}) != nil {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 0, TakeCard: true}) != nil {
		t.Error("Error game move")
	}
	if g.currentPlayerId != 1 {
		t.Error("Error game move")
	}
	if g.Move(&Move{PlayerId: 1}) == nil {
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
	if info.Players[0].HandCount != 2 || info.Players[1].HandCount != 1 {
		t.Error("Error game info")
	}
	if info.Players[0].Points != 0 || info.Players[1].Points != 0 {
		t.Error("Error game info")
	}
	if info.CurrentPlayerId != 1 {
		t.Error("Error game info")
	}
	if info.BancCount != 3 {
		t.Error("Error game info")
	}
	if info.DiscardCount != 0 {
		t.Error("Error game info")
	}
	if info.Hand[0].Code != "8-hearts" {
		t.Error("Error game info")
	}
	if info.IsCompleted {
		t.Error("Error game info")
	}
	if !info.CanMove {
		t.Error("Error game info")
	}
	if info.CanTake {
		t.Error("Error game info")
	}
	if info.CanSkip {
		t.Error("Error game info")
	}
	g.isCompleted = true
	info = g.Info(0)
	if info == nil {
		t.Error("Error game info")
	}
	if !info.IsCompleted {
		t.Error("Error game info")
	}
	if info.Players[0].Points != 19 || info.Players[1].Points != 8 {
		t.Error("Error game info")
	}
}
