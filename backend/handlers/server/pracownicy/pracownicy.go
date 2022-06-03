package pracownicy

import (
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/tab-projekt-backend/auth_middleware"
	"github.com/tab-projekt-backend/database/redis"
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
	get.HandleFunc("", p.getAll)
	get.HandleFunc("/pesel/{pesel:[0-9]{11}}", p.getByPesel)
	get.HandleFunc("/{id:[0-9]+}", p.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", p.createNew)

	r.Use(auth_middleware.NewAuthorisationMiddleware(p.l, auth_middleware.Authorizer{Level: redis.AdministratorDB}).Middleware)
}
