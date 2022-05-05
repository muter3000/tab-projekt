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
	rw.Header().Add("Content-Type", "application/json")
	k.l.Debug("handling post request", "path", k.path)
	data := postData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Creating new kierowca", http.StatusInternalServerError)
		return
	}
	kierowca := schemas.Kierowca{Pracownik: schemas.Pracownik{Pesel: data.Pesel, Imie: data.Imie,
		Nazwisko: data.Nazwisko, Login: data.Login, Haslo: data.Haslo}}

	_, err = k.db.Model(&kierowca).Returning("*", &kierowca).Insert()
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	for _, element := range data.Kategorie {
		kategoria_kierowcy := schemas.KategoriaKierowcy{KierowcaId: kierowca.KierowcaID, KategoriaPrawaJazdyId: element}
		k.db.Model(&kategoria_kierowcy).Insert()
	}

	err = json.NewEncoder(rw).Encode(kierowca)
	if err != nil {
		http.Error(rw, "Error marshaling new pracownik", http.StatusBadRequest)
		return
	}

}
