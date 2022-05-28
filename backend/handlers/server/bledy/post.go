package bledy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
)

func (b *Bledy) createNew(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	b.l.Debug("handling post request", "path", b.path)
	blad := schemas.Blad{}
	err := json.NewDecoder(r.Body).Decode(&blad)
	if blad.Tytul == "" || blad.Opis == "" {
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	if err != nil {
		b.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new blad", http.StatusInternalServerError)
		return
	}
	_, err = b.db.Model(&blad).Returning("*", &blad).Insert()
	if err != nil {
		b.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(blad)
	if err != nil {
		http.Error(rw, "Error marshaling new blad", http.StatusBadRequest)
		return
	}

}
