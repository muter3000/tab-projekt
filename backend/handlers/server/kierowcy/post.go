package kierowcy

import (
	"encoding/json"
	"net/http"

	"github.com/tab-projekt-backend/schemas"
)

func (k *Kierowcy) createNew(rw http.ResponseWriter, r *http.Request) {
	type postData struct {
		Pesel     string  `json:"pesel"`
		Imie      string  `json:"imie"`
		Nazwisko  string  `json:"nazwisko"`
		Login     string  `json:"login"`
		Haslo     string  `json:"haslo,omitempty"`
		Kategorie []int32 `json:"kategorie"`
	}

	tx, err := k.db.Begin()
	// Make sure to close transaction if something goes wrong.
	defer tx.Close()

	rw.Header().Add("Content-Type", "application/json")
	k.l.Debug("handling post request", "path", k.path)
	data := postData{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new kierowca", http.StatusInternalServerError)
		return
	}
	kierowca := schemas.Kierowca{Pracownik: schemas.Pracownik{Pesel: data.Pesel, Imie: data.Imie,
		Nazwisko: data.Nazwisko, Login: data.Login, Haslo: data.Haslo}}

	_, err = tx.Model(&kierowca).Returning("*", &kierowca).Insert()
	if err != nil {
		k.l.Error("inserting kierowca", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		_ = tx.Rollback()
		return
	}

	for _, element := range data.Kategorie {
		kategoria_kierowcy := schemas.KategoriaKierowcy{KierowcaId: kierowca.KierowcaID, KategoriaPrawaJazdyId: element}
		_, err = tx.Model(&kategoria_kierowcy).Insert()

		if err != nil {
			k.l.Error("inserting kategorie", "err", err)
			http.Error(rw, "Inserting into database", http.StatusInternalServerError)
			_ = tx.Rollback()
			return
		}
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	err = json.NewEncoder(rw).Encode(kierowca)
	if err != nil {
		http.Error(rw, "Error marshaling new pracownik", http.StatusBadRequest)
		return
	}

}
