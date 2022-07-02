package bledy

import (
	"net/http"

	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/middlewares"

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
	rGetAndDelete := router.PathPrefix(b.path).Subrouter()
	get := rGetAndDelete.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", b.getAll)
	get.HandleFunc("/{id:[0-9]+}", b.getByID)

	rPost := router.PathPrefix(b.path).Subrouter()
	post := rPost.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", b.createNew)

	delete := rGetAndDelete.Methods(http.MethodDelete).Subrouter()
	delete.HandleFunc("/{id:[0-9]+}", b.deleteByID)

	rPost.Use(middlewares.NewAuthorisationMiddleware(b.l, middlewares.Authorizer{Level: redis.Kierowca}).Middleware)
	rGetAndDelete.Use(middlewares.NewAuthorisationMiddleware(b.l, middlewares.Authorizer{Level: redis.AdministratorDB}).Middleware)
}
