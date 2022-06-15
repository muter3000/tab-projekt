package kierowcy

import (
	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/middlewares"
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

	getKierowca := r.Methods(http.MethodGet).Subrouter()
	getKierowca.HandleFunc("/me", k.getMe)
	getKierowca.Use(middlewares.NewAuthorisationMiddleware(k.l, middlewares.Authorizer{Level: redis.Kierowca}).Middleware)

	getAdmin := r.Methods(http.MethodGet).Subrouter()
	getAdmin.HandleFunc("/pesel/{pesel:[0-9]{11}}", k.getByPesel)
	getAdmin.HandleFunc("/{id:[0-9]+}", k.getByID)
	getAdmin.HandleFunc("", k.getAll)
	getAdmin.Use(middlewares.NewAuthorisationMiddleware(k.l, middlewares.Authorizer{Level: redis.Administrator}).Middleware)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", k.createNew)
	post.Use(middlewares.NewAuthorisationMiddleware(k.l, middlewares.Authorizer{Level: redis.Administrator}).Middleware)
}
