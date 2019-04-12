package main

import (
	"card-games/decks/classic36"
	"card-games/games/pharaoh"
	"encoding/json"
	"log"
)

type RoomPharaoh struct {
	game       *pharaoh.Game
	path       string
	hall       *Hall
	users      map[*User]int
	action     chan *Action
	register   chan *User
	unregister chan *User
}

func newRoomPharaoh(h *Hall, path string) Room {
	return &RoomPharaoh{
		game:       pharaoh.NewGame(),
		path:       path,
		hall:       h,
		action:     make(chan *Action),
		register:   make(chan *User),
		unregister: make(chan *User),
		users:      make(map[*User]int),
	}
}

func (r RoomPharaoh) getPath() string {
	return r.path
}

func (r RoomPharaoh) enter(u *User) error {
	r.register <- u
	return nil
}

func (r RoomPharaoh) leave(u *User) error {
	r.unregister <- u
	return nil
}

func (r RoomPharaoh) close() {
	close(r.register)
	close(r.unregister)
	close(r.action)
}

func (r RoomPharaoh) runAction(a *Action) {
	r.action <- a
}

func (r RoomPharaoh) startGame() {
	var dec = classic36.Deck()
	var cards = make([]pharaoh.Card, len(dec))
	for i, v := range dec {
		cards[i] = pharaoh.Card(v)
	}
	r.game.Start(cards)
	r.hall.play <- &r
}

func (r RoomPharaoh) broadcastInfo() {
	for u, id := range r.users {
		info := r.game.Info(id)
		msg, _ := json.Marshal(info)
		for c := range u.clients {
			select {
			case c.send <- msg:
			default:
				close(c.send)
				delete(u.clients, c)
			}
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
			if len(r.users) == r.game.MaxPlayers {
				r.startGame()
				r.broadcastInfo()
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
