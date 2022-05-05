package kursy

import (
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Kursy struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewKursy(l hclog.Logger, db *pg.DB, path string) *Kursy {
	return &Kursy{l: l, db: db, path: path}
}

func (k *Kursy) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(k.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", k.getAll)
	get.HandleFunc("/{id:[0-9]+}", k.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", k.createNew)
}
