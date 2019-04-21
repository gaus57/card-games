package main

import (
	"encoding/json"
	"log"
	"time"
)

type User struct {
	uid        string
	name       string
	clients    map[*Client]bool
	hall       *Hall
	room       RoomInterface
	closeTimer *time.Timer
}

func newUser(h *Hall, uid string) *User {
	u := &User{
		uid:     uid,
		hall:    h,
		clients: make(map[*Client]bool),
	}
	h.register <- u

	return u
}

func (u *User) close() {
	for c := range u.clients {
		close(c.send)
		delete(u.clients, c)
		c.conn.Close()
	}
	u.leave()
}

func (u *User) startClose() {
	log.Println("close timer", time.Second*10)
	u.closeTimer = time.NewTimer(time.Second * 10)
	go func() {
		select {
		case <-u.closeTimer.C:
			u.hall.unregister <- u
		default:
		}
	}()
}

func (u *User) addClient(c *Client) {
	if u.closeTimer != nil {
		log.Println("close timer stop")
		u.closeTimer.Stop()
	}
	u.clients[c] = true
	if u.room != nil {
		time.Sleep(time.Millisecond * 100)
		u.room.sendInfo(u)
	}
}

func (u *User) closeClient(c *Client) {
	if _, ok := u.clients[c]; ok {
		delete(u.clients, c)
		close(c.send)
		c.conn.Close()
	}
	if len(u.clients) == 0 {
		u.startClose()
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

func (u *User) send(msg Message) {
	text, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error Marshal Message:", err)
		return
	}
	for c := range u.clients {
		select {
		case c.send <- text:
		default:
			close(c.send)
			delete(u.clients, c)
		}
	}
}
