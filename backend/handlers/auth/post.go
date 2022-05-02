package auth

import (
	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/database/redis"
	"net/http"
	"strconv"
)

func (a *AuthHandler) CreateSession(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	level, err := strconv.Atoi(vars["level"])
	if err != nil {
		panic(err)
	}

	res := a.rc.CreateUserSession(rw, redis.PermissionLevel(level))
	if res != true {
		http.Error(rw, "Error creating session", http.StatusInternalServerError)
	}
	rw.WriteHeader(http.StatusCreated)
}
