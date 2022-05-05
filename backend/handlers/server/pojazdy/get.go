package pojazdy

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (p *Pojazdy) getAll(rw http.ResponseWriter, _ *http.Request) {
	p.l.Debug("handling get all request", "path", p.path)

	rw.Header().Add("Content-Type", "application/json")

	var pojazd []schemas.Pojazd
	err := p.db.Model(&pojazd).Select()
	if err != nil {
		p.l.Error("while handling get all", "path", p.path, "error", err)
		http.Error(rw, "Error getting pojazd table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(pojazd)
	if err != nil {
		p.l.Error("err", "", err)
		return
	}
}

func (p *Pojazdy) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	p.l.Debug("handling get by ID request", "path", p.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	pojazd := schemas.Pojazd{}
	err = p.db.Model(&pojazd).Where("id = ?", id).Select()
	if err != nil {
		p.l.Error("while handling get by ID", "path", p.path, "error", err)
		http.Error(rw, "Error getting pojazd table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(pojazd)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
