package kierowcy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
	"golang.org/x/crypto/bcrypt"
)

func (k *Kierowcy) createNew(rw http.ResponseWriter, r *http.Request) {
	k.l.Debug("handling post request", "path", k.path)
	kierowca := schemas.Kierowca{}
	err := json.NewDecoder(r.Body).Decode(&kierowca)
	if kierowca.Haslo == "" || kierowca.Login == "" || kierowca.Imie == "" || kierowca.Nazwisko == "" ||
		kierowca.Pesel == "" || len(kierowca.Pesel) != 11 {
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new kierowca", http.StatusInternalServerError)
	}
	saltedPassword, err := bcrypt.GenerateFromPassword([]byte(kierowca.Haslo), bcrypt.DefaultCost)
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Salting password", http.StatusInternalServerError)
	}
	kierowca.Haslo = string(saltedPassword)
	_, err = k.db.Model(&kierowca).Returning("*", &kierowca).Insert()
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
	}
	retured, err := json.Marshal(kierowca)
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Marshalling response to json", http.StatusInternalServerError)
	}
	rw.Write(retured)
}
