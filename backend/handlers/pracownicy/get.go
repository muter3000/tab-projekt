package pracownicy

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
	"net/http"
	"strconv"
)

func (p *Pracownicy) getAll(rw http.ResponseWriter, _ *http.Request) {
	p.l.Debug("handling get all request for /pracownicy")
	var pracownicy []schemas.Pracownik
	err := p.db.Model(&pracownicy).Select()
	if err != nil {
		http.Error(rw, "Getting pracownicy table", http.StatusInternalServerError)
	}
	err = json.NewEncoder(rw).Encode(pracownicy)
	if err != nil {
		http.Error(rw, "Encoding to json", http.StatusInternalServerError)
	}
	p.l.Debug("")
}

func (p *Pracownicy) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	p.l.Debug("handling get by ID request for /pracownicy", "id", id)

	pracownik := schemas.Pracownik{Id: int32(id)}
	err = p.db.Model(&pracownik).Select()
	if err != nil {
		http.Error(rw, "Getting pracownicy table", http.StatusInternalServerError)
	}
	err = json.NewEncoder(rw).Encode(pracownik)
	if err != nil {
		http.Error(rw, "Encoding to json", http.StatusInternalServerError)
	}
}
