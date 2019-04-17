package main

import "log"

type Hall struct {
	openRooms  map[string]RoomInterface
	playRooms  map[RoomInterface]bool
	users      map[string]*User
	register   chan *User
	unregister chan *User
	enter      chan *Invite
	play       chan RoomInterface
	close      chan RoomInterface
}

type Invite struct {
	path string
	user *User
}

func newHall() *Hall {
	return &Hall{
		openRooms:  make(map[string]RoomInterface),
		playRooms:  make(map[RoomInterface]bool),
		users:      make(map[string]*User),
		register:   make(chan *User),
		unregister: make(chan *User),
		enter:      make(chan *Invite),
		play:       make(chan RoomInterface),
		close:      make(chan RoomInterface),
	}
}

func (h *Hall) getUser(uid string) *User {
	if u, ok := h.users[uid]; ok {
		return u
	}
	return nil
}

func (h *Hall) run() {
	for {
		select {
		case u := <-h.register:
			log.Println("register user", u.uid)
			h.users[u.uid] = u
		case u := <-h.unregister:
			log.Println("unregister user", u.uid)
			if _, ok := h.users[u.uid]; ok {
				delete(h.users, u.uid)
				u.close()
			}
		case invite := <-h.enter:
			log.Println("invite room ", invite.path)
			if invite.user.room != nil && invite.user.room.getPath() == invite.path {
				log.Println("become ", invite.path)
			} else if r, ok := h.openRooms[invite.path]; ok {
				r.enter(invite.user)
			} else {
				log.Println("create room ", invite.path)
				switch invite.path {
				case "/pharaoh":
					r := newRoomPharaoh(h, invite.path)
					h.openRooms[r.getPath()] = r
					go r.run()
					r.enter(invite.user)
				default:
					log.Println("create invalid room", invite.path)
				}
			}
		case r := <-h.close:
			log.Println("close room ", r.getPath())
			if _, ok := h.playRooms[r]; ok {
				delete(h.playRooms, r)
				r.close()
			}
		case r := <-h.play:
			log.Println("play room ", r.getPath())
			if _, ok := h.openRooms[r.getPath()]; ok {
				h.playRooms[r] = true
				delete(h.openRooms, r.getPath())
			}
		}
	}
}
