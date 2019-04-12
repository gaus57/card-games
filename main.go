package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	// "github.com/google/uuid"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// var uid string
	// if cookie, err := r.Cookie("uid"); err == nil {
	// 	uid = cookie.Value
	// }
	// if uid == "" {
	// 	cookie := &http.Cookie{
	// 		Name:    "uid",
	// 		Value:   uuid.New().String(),
	// 		Expires: time.Now().Add(365 * 24 * time.Hour),
	// 	}
	// 	http.SetCookie(w, cookie)
	// }

	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, "pages/index.html")
	case "/pharaoh":
		http.ServeFile(w, r, "pages/pharaoh.html")
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func serveWs(h *Hall, w http.ResponseWriter, r *http.Request) {
	var uid string
	if cookie, err := r.Cookie("uid"); err == nil {
		uid = cookie.Value
	}
	user := h.getUser(uid)
	if user == nil {
		user = newUser(h)
		cookie := &http.Cookie{
			Name:    "uid",
			Value:   user.uid,
			Expires: time.Now().Add(365 * 24 * time.Hour),
		}
		http.SetCookie(w, cookie)
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{user: user, conn: conn, send: make(chan []byte)}
	user.addClient(client)

	go client.writePump()
	go client.readPump()
}

func main() {
	flag.Parse()
	hall := newHall()
	go hall.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hall, w, r)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
