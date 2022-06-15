package stanowisko_administracyjne

import (
	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/middlewares"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Stanowisko_administracyjne struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewStanowiskoAdministracyjne(l hclog.Logger, db *pg.DB, path string) *Stanowisko_administracyjne {
	return &Stanowisko_administracyjne{l: l, db: db, path: path}
}

func (sa *Stanowisko_administracyjne) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(sa.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", sa.getAll)
	get.HandleFunc("/{id:[0-9]+}", sa.getByID)

	get.Use(middlewares.NewAuthorisationMiddleware(sa.l, middlewares.Authorizer{Level: redis.Administrator}).Middleware)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", sa.createNew)

	post.Use(middlewares.NewAuthorisationMiddleware(sa.l, middlewares.Authorizer{Level: redis.AdministratorDB}).Middleware)
}
