package pracownicy

import (
	"encoding/json"
	"github.com/tab-projekt-backend/schemas"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func (p *Pracownicy) createNew(rw http.ResponseWriter, r *http.Request) {
	p.l.Debug("handling post request", "path", p.path)
	pracownik := schemas.Pracownik{}
	err := json.NewDecoder(r.Body).Decode(&pracownik)
	if err != nil {
		p.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new pracownik", http.StatusBadRequest)
	}
	saltedPassword, err := bcrypt.GenerateFromPassword([]byte(pracownik.Haslo), bcrypt.DefaultCost)
	if err != nil {
		p.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new pracownik", http.StatusBadRequest)
	}
	pracownik.Haslo = string(saltedPassword)
	_, err = p.db.Model(&pracownik).Insert()
	if err != nil {
		p.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new pracownik", http.StatusBadRequest)
	}
	log.New(rw, "server", 0).Println("Success")

}
