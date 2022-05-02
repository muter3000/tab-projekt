package trasy

import (
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Trasy struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewTrasy(l hclog.Logger, db *pg.DB, path string) *Trasy {
	return &Trasy{l: l, db: db, path: path}
}

func (t *Trasy) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(t.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("/", t.getAll)
	get.HandleFunc("/{id:[0-9]+}", t.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/", t.createNew)
}
