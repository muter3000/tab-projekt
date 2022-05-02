package marki

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
)

func (m *Marki) createNew(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	m.l.Debug("handling post request", "path", m.path)
	marka := schemas.Marka{}
	err := json.NewDecoder(r.Body).Decode(&marka)
	if marka.Nazwa == "" {
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	if err != nil {
		m.l.Error("marshaling", "err", err)
		http.Error(rw, "Error creating new marka", http.StatusBadRequest)
		return
	}
	_, err = m.db.Model(&marka).Returning("*", &marka).Insert()
	if err != nil {
		m.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(marka)
	if err != nil {
		http.Error(rw, "Error marshaling new marka", http.StatusBadRequest)
		return
	}

}
