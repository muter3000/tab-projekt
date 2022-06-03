package marki

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (m *Marki) getAll(rw http.ResponseWriter, _ *http.Request) {
	m.l.Debug("handling get all request", "path", m.path)

	rw.Header().Add("Content-Type", "application/json")

	var marka []schemas.Marka
	err := m.db.Model(&marka).Select()
	if err != nil {
		m.l.Error("while handling get all", "path", m.path, "error", err)
		http.Error(rw, "Error getting marka table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(marka)
	if err != nil {
		m.l.Error("err", "", err)
		return
	}
}

func (m *Marki) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	m.l.Debug("handling get by ID request", "path", m.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	marka := schemas.Pojazd{}
	err = m.db.Model(&marka).Where("id = ?", id).Select()
	if err != nil {
		m.l.Error("while handling get by ID", "path", m.path, "error", err)
		http.Error(rw, "Error getting marka table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(marka)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
