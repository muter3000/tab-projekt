package administratorzy

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (a *Administratorzy) getAll(rw http.ResponseWriter, _ *http.Request) {
	a.l.Debug("handling get all request", "path", a.path)

	rw.Header().Add("Content-Type", "application/json")

	var administracja []schemas.Administrator
	err := a.db.Model(&administracja).Relation("StanowiskoAdministracyjne").Select()
	if err != nil {
		a.l.Error("while handling get all", "path", a.path, "error", err)
		http.Error(rw, "Error getting administracja table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(administracja)
	if err != nil {
		a.l.Error("err", "", err)
		http.Error(rw, "Error encoding administracja table to JSON", http.StatusInternalServerError)
		return
	}
}

func (a *Administratorzy) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	a.l.Debug("handling get by ID request", "path", a.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	administrator := schemas.Administrator{}
	err = a.db.Model(&administrator).Where("administracja_id = ?", id).Relation("StanowiskoAdministracyjne").Select()
	if err != nil {
		a.l.Error("while handling get by ID", "path", a.path, "error", err)
		http.Error(rw, "Error getting administracja table", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(administrator)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}

func (a *Administratorzy) getByPesel(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	pesel := vars["pesel"]
	a.l.Debug("handling get by pesel request", "path", a.path, "pesel", pesel)

	rw.Header().Add("Content-Type", "application/json")

	administrator := schemas.Administrator{}
	err := a.db.Model(&administrator).Where("pesel = ?", pesel).Select()
	if err != nil {
		a.l.Error("while handling get by ID", "path", a.path, "error", err)
		http.Error(rw, "Error getting administracja table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(administrator)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
