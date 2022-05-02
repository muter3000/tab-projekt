package auth

import (
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/tab-projekt-backend/database/redis"
	"net/http"
)

type AuthHandler struct {
	l    hclog.Logger
	rc   *redis.RedisClient
	pg   *pg.DB
	path string
}

func NewAuthHandler(rc *redis.RedisClient) *AuthHandler {
	return &AuthHandler{rc: rc}
}

func (a *AuthHandler) RegisterSubRouter(router *mux.Router) {
	r := router.PathPrefix(a.path).Subrouter()
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("/{level:[0-2]{1}}", a.CheckAuthorization)

	post := r.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/{level:[0-2]{1}}", a.CreateSession)

	del := r.Methods(http.MethodDelete).Subrouter()
	del.HandleFunc("/", a.InvalidateSession)
}
