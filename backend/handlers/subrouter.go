package handlers

import "github.com/gorilla/mux"

type SubRouter interface {
	// RegisterSubRouter Funkcja rejestruje subrouter dla podanego routera
	RegisterSubRouter(router *mux.Router)
}
