package pojazdy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
)

func (p *Pojazdy) createNew(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	p.l.Debug("handling post request", "path", p.path)
	pojazd := schemas.Pojazd{}
	err := json.NewDecoder(r.Body).Decode(&pojazd)
	if pojazd.NumerRejestracyjny == "" || err != nil {
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	_, err = p.db.Model(&pojazd).Returning("*", &pojazd).Insert()
	if err != nil {
		p.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(pojazd)
	if err != nil {
		http.Error(rw, "Error marshaling new pojazd", http.StatusBadRequest)
		return
	}

}
