package trasy

import (
<<<<<<< HEAD
	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/middlewares"
=======
	"github.com/tab-projekt-backend/auth_middleware"
	"github.com/tab-projekt-backend/database/redis"
>>>>>>> f5cb11a608279d707ff2189eaa32a6ada4ad0931
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
	get.HandleFunc("", t.getAll)
	get.HandleFunc("/{id:[0-9]+}", t.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", t.createNew)

<<<<<<< HEAD
	r.Use(middlewares.NewAuthorisationMiddleware(t.l, middlewares.Authorizer{Level: redis.Administrator}).Middleware)
=======
	r.Use(auth_middleware.NewAuthorisationMiddleware(t.l, auth_middleware.Authorizer{Level: redis.Administrator}).Middleware)
>>>>>>> f5cb11a608279d707ff2189eaa32a6ada4ad0931
}
