package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

const (
	SessionKey  = "session-key"
	SessionName = "session-name"
)

func CheckSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SessionName)

	user, ok := session.Values[SessionKey].(string)
	if !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintf(w, "Authenticated user: %v", user)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SessionName)
	session.Values[SessionKey] = "email"
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/check", CheckSession)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
