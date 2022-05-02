package auth

import (
	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/database/redis"
	"net/http"
	"strconv"
)

func (a *AuthHandler) CheckAuthorization(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	level, err := strconv.Atoi(vars["level"])
	if err != nil {
		panic(err)
	}

	auth := a.rc.CheckAuthorization(r, redis.PermissionLevel(level))
	if auth == false {
		http.Error(rw, "Error: not authorized", http.StatusUnauthorized)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
