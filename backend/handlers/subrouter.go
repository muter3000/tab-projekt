package handlers

import "github.com/gorilla/mux"

type SubRouter interface {
	RegisterSubRouter(router *mux.Router)
}
