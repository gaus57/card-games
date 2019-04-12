package main

import (
	"github.com/google/uuid"
)

type User struct {
	uid      string
	playerId int
	clients  map[*Client]bool
	hall     *Hall
	room     Room
}

func newUser(h *Hall) *User {
	u := &User{
		uid:     uuid.New().String(),
		hall:    h,
		clients: make(map[*Client]bool),
	}
	h.register <- u

	return u
}

func (u *User) addClient(c *Client) {
	u.clients[c] = true
}

func (u *User) closeClient(c *Client) {
	if _, ok := u.clients[c]; ok {
		delete(u.clients, c)
	}
	close(c.send)
	c.conn.Close()
	if len(u.clients) == 0 {
		u.leave()
		u.hall.unregister <- u
	}
}

func (u *User) enter(path string) {
	u.hall.enter <- &Invite{path: path, user: u}
}

func (u *User) leave() {
	if u.room != nil {
		u.room.leave(u)
	}
}
