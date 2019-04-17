package main

type RoomInterface interface {
	getPath() string
	enter(*User) error
	leave(*User) error
	runAction(*Action)
	run()
	close()
	sendInfo(*User)
}

type Action struct {
	user *User
	text []byte
}

type Message struct {
	Event string
	Data  interface{}
}

type Room struct {
	path       string
	hall       *Hall
	users      map[*User]int
	action     chan *Action
	register   chan *User
	unregister chan *User
}

func newRoom(h *Hall, path string) *Room {
	return &Room{
		path:       path,
		hall:       h,
		action:     make(chan *Action),
		register:   make(chan *User),
		unregister: make(chan *User),
		users:      make(map[*User]int),
	}
}

func (r Room) getPath() string {
	return r.path
}

func (r Room) enter(u *User) error {
	r.register <- u
	return nil
}

func (r Room) leave(u *User) error {
	r.unregister <- u
	return nil
}

func (r Room) close() {
	close(r.register)
	close(r.unregister)
	close(r.action)
}

func (r Room) runAction(a *Action) {
	r.action <- a
}

func (r Room) broadcast(msg Message) {
	for u, _ := range r.users {
		u.send(msg)
	}
}
