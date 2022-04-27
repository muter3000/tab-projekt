package pracownicy

import (
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Pracownicy struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewPracownicy(l hclog.Logger, db *pg.DB, path string) *Pracownicy {
	return &Pracownicy{l: l, db: db, path: path}
}

func (p *Pracownicy) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(p.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("/", p.getAll)
	get.HandleFunc("/{id:[0-9]+}", p.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/", p.createNew)
}
