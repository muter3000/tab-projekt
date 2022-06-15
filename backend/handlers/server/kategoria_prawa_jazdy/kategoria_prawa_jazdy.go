package kategoria_prawa_jazdy

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

type Kategoria_prawa_jazdy struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewKategoriaPrawaJazdy(l hclog.Logger, db *pg.DB, path string) *Kategoria_prawa_jazdy {
	return &Kategoria_prawa_jazdy{l: l, db: db, path: path}
}

func (kpj *Kategoria_prawa_jazdy) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(kpj.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", kpj.getAll)
	get.HandleFunc("/{id:[0-9]+}", kpj.getByID)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", kpj.createNew)
<<<<<<< HEAD
	r.Use(middlewares.NewAuthorisationMiddleware(kpj.l, middlewares.Authorizer{Level: redis.AdministratorDB}).Middleware)
=======
	r.Use(auth_middleware.NewAuthorisationMiddleware(kpj.l, auth_middleware.Authorizer{Level: redis.AdministratorDB}).Middleware)
>>>>>>> f5cb11a608279d707ff2189eaa32a6ada4ad0931
}
