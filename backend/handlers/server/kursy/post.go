package kursy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
)

func (k *Kursy) createNew(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	k.l.Debug("handling post request", "path", k.path)
	kurs := schemas.Kurs{}
	err := json.NewDecoder(r.Body).Decode(&kurs)
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	_, err = k.db.Model(&kurs).Relation("Trasa").Relation("Kierowca").Relation("Pojazd").Relation("Pojazd.Marka").Relation("Kierowca.Kategorie").Returning("*", &kurs).Insert()
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(kurs)
	if err != nil {
		http.Error(rw, "Error marshaling new kurs", http.StatusBadRequest)
		return
	}

}
