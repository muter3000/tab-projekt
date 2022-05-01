package stanowisko_administracyjne

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
)

func (sa *Stanowisko_administracyjne) createNew(rw http.ResponseWriter, r *http.Request) {
	sa.l.Debug("handling post request", "path", sa.path)
	stanowisko := schemas.StanowiskoAdministracyjne{}
	err := json.NewDecoder(r.Body).Decode(&stanowisko)

	if err != nil {
		sa.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new stanowisko", http.StatusInternalServerError)
		return
	}
	_, err = sa.db.Model(&stanowisko).Returning("*", &stanowisko).Insert()
	if err != nil {
		sa.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(stanowisko)
	if err != nil {
		http.Error(rw, "Error marshaling new stanowisko_administracyjne", http.StatusBadRequest)
		return
	}

}
