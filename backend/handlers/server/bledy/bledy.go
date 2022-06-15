package bledy

import (
	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/middlewares"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Bledy struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewBledy(l hclog.Logger, db *pg.DB, path string) *Bledy {
	return &Bledy{l: l, db: db, path: path}
}

func (b *Bledy) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(b.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", b.getAll)
	get.HandleFunc("/{id:[0-9]+}", b.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", b.createNew)

	r.Use(middlewares.NewAuthorisationMiddleware(b.l, middlewares.Authorizer{Level: redis.Kierowca}).Middleware)
}
