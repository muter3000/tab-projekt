package kierowcy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
	"golang.org/x/crypto/bcrypt"
)

func (k *Kierowcy) createNew(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	k.l.Debug("handling post request", "path", k.path)
	kierowca := schemas.Kierowca{}
	err := json.NewDecoder(r.Body).Decode(&kierowca)

	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new kierowca", http.StatusInternalServerError)
		return
	}
	saltedPassword, err := bcrypt.GenerateFromPassword([]byte(kierowca.Haslo), bcrypt.DefaultCost)
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Salting password", http.StatusInternalServerError)
		return
	}
	kierowca.Haslo = string(saltedPassword)
	_, err = k.db.Model(&kierowca).Returning("*", &kierowca).Insert()
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	kierowca.Haslo = ""

	err = json.NewEncoder(rw).Encode(kierowca)
	if err != nil {
		http.Error(rw, "Error marshaling new pracownik", http.StatusBadRequest)
		return
	}

}
