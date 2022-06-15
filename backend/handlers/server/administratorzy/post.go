package administratorzy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
)

func (a *Administratorzy) createNew(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	a.l.Debug("handling post request", "path", a.path)
	administrator := schemas.Administrator{}
	err := json.NewDecoder(r.Body).Decode(&administrator)
	if administrator.Haslo == "" || administrator.Login == "" || administrator.Imie == "" || administrator.Nazwisko == "" ||
		administrator.Pesel == "" || len(administrator.Pesel) != 11 {
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	if err != nil {
		a.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new administrator", http.StatusInternalServerError)
		return
	}

	_, err = a.db.Model(&administrator).Returning("*", &administrator).Insert()
	if err != nil {
		a.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	administrator.Haslo = ""

	err = json.NewEncoder(rw).Encode(administrator)
	if err != nil {
		http.Error(rw, "Error marshaling new pracownik", http.StatusBadRequest)
		return
	}

}
