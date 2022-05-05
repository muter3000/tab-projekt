package auth

import "net/http"

func (a *AuthHandler) InvalidateSession(rw http.ResponseWriter, r *http.Request) {
	res := a.rc.InvalidateUserSession(r)
	if res != true {
		http.Error(rw, "Couldn't invalidate session (maybe user wasn't signed in?)", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
