package trasy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
)

func (t *Trasy) createNew(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	t.l.Debug("handling post request", "path", t.path)
	trasa := schemas.Trasa{}
	err := json.NewDecoder(r.Body).Decode(&trasa)
	if trasa.MiejscowoscPoczatkowa == "" || trasa.MiejscowoscKoncowa == "" {
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	if err != nil {
		t.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new trasa", http.StatusInternalServerError)
		return
	}
	_, err = t.db.Model(&trasa).Returning("*", &trasa).Insert()
	if err != nil {
		t.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(trasa)
	if err != nil {
		http.Error(rw, "Error marshaling new trasa", http.StatusBadRequest)
		return
	}

}
