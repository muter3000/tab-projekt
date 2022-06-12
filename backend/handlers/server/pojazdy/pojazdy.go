package pojazdy

import (
	"github.com/tab-projekt-backend/auth_middleware"
	"github.com/tab-projekt-backend/database/redis"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Pojazdy struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewPojazdy(l hclog.Logger, db *pg.DB, path string) *Pojazdy {
	return &Pojazdy{l: l, db: db, path: path}
}

func (p *Pojazdy) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(p.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", p.getAll)
	get.HandleFunc("/{id:[0-9]+}", p.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", p.createNew)
	r.Use(auth_middleware.NewAuthorisationMiddleware(p.l, auth_middleware.Authorizer{Level: redis.Administrator}).Middleware)
}
