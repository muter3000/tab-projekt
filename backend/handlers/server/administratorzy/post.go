package administratorzy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
	"golang.org/x/crypto/bcrypt"
)

func (a *Administratorzy) createNew(rw http.ResponseWriter, r *http.Request) {
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
		http.Error(rw, "Error creating new administrator", http.StatusInternalServerError)
		return
	}
	saltedPassword, err := bcrypt.GenerateFromPassword([]byte(administrator.Haslo), bcrypt.DefaultCost)
	if err != nil {
		a.l.Error("marshaling", "err", err)
		http.Error(rw, "Error salting password", http.StatusInternalServerError)
		return
	}
	administrator.Haslo = string(saltedPassword)
	_, err = a.db.Model(&administrator).Returning("*", &administrator).Insert()
	if err != nil {
		a.l.Error("marshaling", "err", err)
		http.Error(rw, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(administrator)
	if err != nil {
		a.l.Error("marshaling", "err", err)
		http.Error(rw, "Error marshalling response to json", http.StatusInternalServerError)
		return
	}
}
