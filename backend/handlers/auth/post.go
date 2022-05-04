package auth

import (
	"encoding/json"
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

	type userCredentials struct {
		Login    string `json:"login"`
		Password string `json:"haslo"`
	}
	userCreds := userCredentials{}

	err = json.NewDecoder(r.Body).Decode(&userCreds)
	if err != nil {
		http.Error(rw, "Error decoding user credentials from request", http.StatusBadRequest)
		return
	}
	if userCreds.Login == "" || userCreds.Password == "" {
		http.Error(rw, "Error getting login and password from body (password or login where blank)", http.StatusBadRequest)
		return
	}

	res := a.ac.CreateUserSession(rw, userCreds.Login, userCreds.Password, redis.PermissionLevel(level))
	if res != true {
		http.Error(rw, "Error creating session", http.StatusUnauthorized)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
