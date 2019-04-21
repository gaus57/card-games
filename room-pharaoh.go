package main

import (
	"card-games/decks/classic36"
	"card-games/games/pharaoh"
	"encoding/json"
	"log"
	"time"
)

type RoomPharaoh struct {
	*Room
	game *pharaoh.Game
}

func newRoomPharaoh(h *Hall, path string) RoomInterface {
	return &RoomPharaoh{
		Room: newRoom(h, path),
		game: pharaoh.NewGame(),
	}
}

func (r RoomPharaoh) startGame() {
	var dec = classic36.Deck()
	var cards = make([]pharaoh.Card, len(dec))
	for i, v := range dec {
		cards[i] = pharaoh.Card(v)
	}
	r.game.Start(cards)
	r.hall.play <- &r
	r.broadcastInfo()
}

func (r RoomPharaoh) broadcastInfo() {
	for u, _ := range r.users {
		r.sendInfo(u)
	}
}

func (r RoomPharaoh) sendInfo(u *User) {
	if !r.game.IsStarted() {
		return
	}
	if id, ok := r.users[u]; ok {
		msg := Message{
			Event: "game-info",
			Data:  r.game.Info(id),
		}
		u.send(msg)
		if r.game.IsCompleted() {
			u.room = nil
		}
	}
}

func (r RoomPharaoh) run() {
	for {
		select {
		case u := <-r.register:
			log.Println("register user in room ", r.path)
			u.room = r
			r.users[u], _ = r.game.Join()
			if r.game.MinPlayers < r.game.MaxPlayers && len(r.users) == r.game.MinPlayers {
				timer := time.AfterFunc(time.Second*10, func() {
					r.startGame()
				})
				defer timer.Stop()
			}
			if len(r.users) == r.game.MaxPlayers {
				r.startGame()
			}
		case u := <-r.unregister:
			log.Println("unregister user in room ", r.path)
			if _, ok := r.users[u]; ok {
				delete(r.users, u)
				u.room = nil
				if len(r.users) == 0 {
					r.hall.close <- &r
				}
			}
		case action := <-r.action:
			log.Println("game action in room ", r.path)
			if playerId, ok := r.users[action.user]; ok {
				var move pharaoh.Move
				err := json.Unmarshal(action.text, &move)
				if err == nil {
					move.PlayerId = playerId
					gameErr := r.game.Move(&move)
					log.Println("game action ", move)
					if gameErr == nil {
						r.broadcastInfo()
						if r.game.IsCompleted() {
							r.hall.close <- &r
						}
					} else {
						log.Println("error action", gameErr)
					}
				} else {
					log.Println("error unmarshal", r.path, err)
				}
			}
		}
	}
}
