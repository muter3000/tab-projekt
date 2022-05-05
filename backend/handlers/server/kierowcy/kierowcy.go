package kierowcy

import (
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Kierowcy struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewKierowcy(l hclog.Logger, db *pg.DB, path string) *Kierowcy {
	return &Kierowcy{l: l, db: db, path: path}
}

func (k *Kierowcy) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(k.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", k.getAll)
	get.HandleFunc("/pesel/{pesel:[0-9]{11}}", k.getByPesel)
	get.HandleFunc("/{id:[0-9]+}", k.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", k.createNew)
}
