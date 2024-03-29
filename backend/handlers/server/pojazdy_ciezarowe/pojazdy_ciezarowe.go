package pojazdy_ciezarowe

import (
	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/middlewares"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type PojazdyCiezarowe struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewPojazdyCiezarowe(l hclog.Logger, db *pg.DB, path string) *PojazdyCiezarowe {
	return &PojazdyCiezarowe{l: l, db: db, path: path}
}

func (pc *PojazdyCiezarowe) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(pc.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", pc.getAll)
	get.HandleFunc("/{id:[0-9]+}", pc.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", pc.createNew)
	r.Use(middlewares.NewAuthorisationMiddleware(pc.l, middlewares.Authorizer{Level: redis.Administrator}).Middleware)
}
