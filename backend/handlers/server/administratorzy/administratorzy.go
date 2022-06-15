package administratorzy

import (
	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/middlewares"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Administratorzy struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewAdministratorzy(l hclog.Logger, db *pg.DB, path string) *Administratorzy {
	return &Administratorzy{l: l, db: db, path: path}
}

func (a *Administratorzy) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(a.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", a.getAll)
	get.HandleFunc("/pesel/{pesel:[0-9]{11}}", a.getByPesel)
	get.HandleFunc("/{id:[0-9]+}", a.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", a.createNew)

	r.Use(middlewares.NewAuthorisationMiddleware(a.l, middlewares.Authorizer{Level: redis.AdministratorDB}).Middleware)
}
