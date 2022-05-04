package auth

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/tab-projekt-backend/database/redis"
	"net/http"
)

type AuthHandler struct {
	l    hclog.Logger
	ac   *redis.AuthorizationClient
	path string
}

func NewAuthHandler(l hclog.Logger, ac *redis.AuthorizationClient, path string) *AuthHandler {
	return &AuthHandler{l: l, ac: ac, path: path}
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
