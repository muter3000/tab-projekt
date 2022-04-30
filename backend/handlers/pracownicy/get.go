package pracownicy

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
	"net/http"
	"strconv"
)

func (p *Pracownicy) getAll(rw http.ResponseWriter, _ *http.Request) {
	p.l.Debug("handling get all request", "path", p.path)

	rw.Header().Add("Content-Type", "application/json")

	var pracownicy []schemas.Pracownik
	err := p.db.Model(&pracownicy).Select()
	if err != nil {
		p.l.Error("while handling get all", "path", p.path, "error", err)
		http.Error(rw, "Error getting pracownicy table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(pracownicy)
	if err != nil {
		p.l.Error("err", "", err)
		return
	}
}

func (p *Pracownicy) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	p.l.Debug("handling get by ID request", "path", p.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	pracownik := schemas.Pracownik{Id: int32(id)}
	err = p.db.Model(&pracownik).Select()
	if err != nil {
		p.l.Error("while handling get by ID", "path", p.path, "error", err)
		http.Error(rw, "Error getting pracownicy table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(pracownik)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}

func (p *Pracownicy) getByPesel(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	pesel := vars["pesel"]
	p.l.Debug("handling get by pesel request", "path", p.path, "pesel", pesel)

	rw.Header().Add("Content-Type", "application/json")

	pracownik := schemas.Pracownik{}
	err := p.db.Model(&pracownik).Where("pesel = ?", pesel).Select()
	if err != nil {
		p.l.Error("while handling get by ID", "path", p.path, "error", err)
		http.Error(rw, "Error getting pracownicy table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(pracownik)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
