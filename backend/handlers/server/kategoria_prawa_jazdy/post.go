package kategoria_prawa_jazdy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
)

func (sa *KategoriaPrawaJazdy) createNew(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	sa.l.Debug("handling post request", "path", sa.path)
	kategoria := schemas.KategoriaPrawaJazdy{}
	err := json.NewDecoder(r.Body).Decode(&kategoria)

	if err != nil {
		sa.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new kategoria", http.StatusInternalServerError)
		return
	}
	_, err = sa.db.Model(&kategoria).Returning("*", &kategoria).Insert()
	if err != nil {
		sa.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(kategoria)
	if err != nil {
		http.Error(rw, "Error marshaling new kategoria prawa jazdy", http.StatusBadRequest)
		return
	}

}
