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

	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Error creating new kierowca", http.StatusInternalServerError)
		return
	}
	saltedPassword, err := bcrypt.GenerateFromPassword([]byte(kierowca.Haslo), bcrypt.DefaultCost)
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Error salting password", http.StatusInternalServerError)
		return
	}
	kierowca.Haslo = string(saltedPassword)
	_, err = k.db.Model(&kierowca).Returning("*", &kierowca).Insert()
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(kierowca)
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Error marshalling response to json", http.StatusInternalServerError)
		return
	}
}
