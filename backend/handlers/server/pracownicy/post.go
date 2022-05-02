package pracownicy

import (
	"encoding/json"

	"github.com/tab-projekt-backend/schemas"
	"net/http"
)

func (p *Pracownicy) createNew(rw http.ResponseWriter, r *http.Request) {

	p.l.Debug("handling post request", "path", p.path)

	pracownik := schemas.Pracownik{}

	err := json.NewDecoder(r.Body).Decode(&pracownik)
	if pracownik.Haslo == "" || pracownik.Login == "" || pracownik.Imie == "" || pracownik.Nazwisko == "" ||
		pracownik.Pesel == "" || len(pracownik.Pesel) != 11 {
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	if err != nil {
		p.l.Error("marshaling", "err", err)
		http.Error(rw, "Error creating new pracownik", http.StatusBadRequest)
		return
	}

	_, err = p.db.Model(&pracownik).Returning("id", &pracownik).Insert()
	if err != nil {
		http.Error(rw, "Error creating new pracownik", http.StatusBadRequest)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(pracownik)
	if err != nil {
		http.Error(rw, "Error marshaling new pracownik", http.StatusBadRequest)
		return
	}

}
