package stanowisko_administracyjne

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (sa *Stanowisko_administracyjne) getAll(rw http.ResponseWriter, _ *http.Request) {
	sa.l.Debug("handling get all request", "path", sa.path)

	rw.Header().Add("Content-Type", "application/json")

	var stanowiska []schemas.StanowiskoAdministracyjne
	err := sa.db.Model(&stanowiska).Select()
	if err != nil {
		sa.l.Error("while handling get all", "path", sa.path, "error", err)
		http.Error(rw, "Error getting stanowiska table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(stanowiska)
	if err != nil {
		sa.l.Error("err", "", err)
		return
	}
}

func (sa *Stanowisko_administracyjne) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	sa.l.Debug("handling get by ID request", "path", sa.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	stanowisko := schemas.StanowiskoAdministracyjne{}
	err = sa.db.Model(&stanowisko).Where("id = ?", id).Select()
	if err != nil {
		sa.l.Error("while handling get by ID", "path", sa.path, "error", err)
		http.Error(rw, "Error getting stanowiska table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(stanowisko)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
