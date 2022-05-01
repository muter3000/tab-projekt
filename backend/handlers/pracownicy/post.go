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
	if err != nil {
		p.l.Error("marshaling", "err", err)
		http.Error(rw, "Error creating new pracownik", http.StatusBadRequest)
	}

	_, err = p.db.Model(&pracownik).Returning("id", &pracownik).Insert()
	if err != nil {
		http.Error(rw, "Error creating new pracownik", http.StatusBadRequest)
	}

	pracownik.Haslo = ""

	err = json.NewEncoder(rw).Encode(pracownik)
	if err != nil {
		http.Error(rw, "Error marshaling new pracownik", http.StatusBadRequest)
	}
}
