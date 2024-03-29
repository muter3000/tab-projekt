package marki

import (
	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/middlewares"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Marki struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewMarki(l hclog.Logger, db *pg.DB, path string) *Marki {
	return &Marki{l: l, db: db, path: path}
}

func (m *Marki) RegisterSubRouter(router *mux.Router) {
	adminMiddleware := middlewares.NewAuthorisationMiddleware(m.l, middlewares.Authorizer{Level: redis.Administrator}).Middleware
	r := router.PathPrefix(m.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", m.getAll)
	get.HandleFunc("/{id:[0-9]+}", m.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", m.createNew)
	r.Use(adminMiddleware)
}
