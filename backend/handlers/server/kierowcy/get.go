package kierowcy

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (k *Kierowcy) getAll(rw http.ResponseWriter, _ *http.Request) {
	k.l.Debug("handling get all request", "path", k.path)

	rw.Header().Add("Content-Type", "application/json")

	var kierowcy []schemas.Kierowca
	err := k.db.Model(&kierowcy).Relation("Kategorie").Select()
	if err != nil {
		k.l.Error("while handling get all", "path", k.path, "error", err)
		http.Error(rw, "Error getting kierowcy table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(kierowcy)
	if err != nil {
		k.l.Error("err", "", err)
		return
	}
}

func (k *Kierowcy) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	k.l.Debug("handling get by ID request", "path", k.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	kierowca := schemas.Kierowca{}
	err = k.db.Model(&kierowca).Where("kierowca_id = ?", id).Relation("Kategorie").Select()
	if err != nil {
		k.l.Error("while handling get by ID", "path", k.path, "error", err)
		http.Error(rw, "Error getting kierowcy table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(kierowca)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}

func (k *Kierowcy) getByPesel(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	pesel := vars["pesel"]
	k.l.Debug("handling get by pesel request", "path", k.path, "pesel", pesel)

	rw.Header().Add("Content-Type", "application/json")

	kierowca := schemas.Kierowca{}
	err := k.db.Model(&kierowca).Relation("Kategorie").Where("pesel = ?", pesel).Select()
	if err != nil {
		k.l.Error("while handling get by ID", "path", k.path, "error", err)
		http.Error(rw, "Error getting kierowcy table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(kierowca)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
