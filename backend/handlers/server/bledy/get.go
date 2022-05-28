package bledy

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (b *Bledy) getAll(rw http.ResponseWriter, _ *http.Request) {
	b.l.Debug("handling get all request", "path", b.path)

	rw.Header().Add("Content-Type", "application/json")

	var blad []schemas.Blad
	err := b.db.Model(&blad).Select()
	if err != nil {
		b.l.Error("while handling get all", "path", b.path, "error", err)
		http.Error(rw, "Error getting blad table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(blad)
	if err != nil {
		b.l.Error("err", "", err)
		return
	}
}

func (b *Bledy) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	b.l.Debug("handling get by ID request", "path", b.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	blad := schemas.Blad{}
	err = b.db.Model(&blad).Where("id = ?", id).Select()
	if err != nil {
		b.l.Error("while handling get by ID", "path", b.path, "error", err)
		http.Error(rw, "Error getting blad table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(blad)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
