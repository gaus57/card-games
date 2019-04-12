package main

type Room interface {
	getPath() string
	enter(*User) error
	leave(*User) error
	runAction(*Action)
	run()
	close()
}

type Action struct {
	user *User
	text []byte
}

type Message struct {
	Event string
	Data  interface{}
}
