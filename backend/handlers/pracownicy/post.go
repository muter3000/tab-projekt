package pracownicy

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
	"golang.org/x/crypto/bcrypt"
)

func (p *Pracownicy) createNew(rw http.ResponseWriter, r *http.Request) {
	p.l.Debug("handling post request", "path", p.path)
	pracownik := schemas.Pracownik{}
	err := json.NewDecoder(r.Body).Decode(&pracownik)
	if pracownik.Haslo == "" || pracownik.Login == "" || pracownik.Imie == "" || pracownik.Nazwisko == "" ||
		pracownik.Pesel == "" || len(pracownik.Pesel) != 11 {
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
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
	_, err = p.db.Model(&pracownik).Returning("id", &pracownik).Insert()
	if err != nil {
		p.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new pracownik", http.StatusBadRequest)
	}
	log.New(rw, "server", 0).Println("Success")

}
