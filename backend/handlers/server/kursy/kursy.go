package kursy

import (
	"github.com/tab-projekt-backend/auth_middleware"
	"github.com/tab-projekt-backend/database/redis"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Kursy struct {
	l    hclog.Logger
	db   *pg.DB
	path string
}

func NewKursy(l hclog.Logger, db *pg.DB, path string) *Kursy {
	return &Kursy{l: l, db: db, path: path}
}

func (k *Kursy) RegisterSubRouter(router *mux.Router) {
	adminMiddleware := auth_middleware.NewAuthorisationMiddleware(k.l, auth_middleware.Authorizer{Level: redis.Kierowca}).Middleware
	kierowcaMiddleware := auth_middleware.NewAuthorisationMiddleware(k.l, auth_middleware.Authorizer{Level: redis.Kierowca}).Middleware

	r := router.PathPrefix(k.path).Subrouter()
	getAdmin := r.Methods(http.MethodGet).Subrouter()
	getAdmin.HandleFunc("", k.getAll)
	getAdmin.HandleFunc("/kierowca/{id:[0-9]+}", k.getByDriverID)
	getAdmin.HandleFunc("/{id:[0-9]+}", k.getByID)
	getAdmin.Use(adminMiddleware)

	getKierowca := r.Methods(http.MethodGet).Subrouter()
	getKierowca.HandleFunc("/me", k.getByMyID)
	getKierowca.Use(kierowcaMiddleware)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", k.createNew)
	post.Use(adminMiddleware)

	patchAdmin := r.Methods(http.MethodPatch).Subrouter()
	patchAdmin.HandleFunc("/{id:[0-9]+}", k.updateExisting)
	patchAdmin.Use(adminMiddleware)

	patchKierowca := r.Methods(http.MethodPatch).Subrouter()
	patchKierowca.HandleFunc("/me", k.updateMyExisting)
	patchKierowca.Use(kierowcaMiddleware)
}
