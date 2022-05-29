package kursy

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (k *Kursy) getAll(rw http.ResponseWriter, r *http.Request) {
	k.l.Debug("handling get all request", "path", k.path)
	query := r.URL.Query()
	var dataRozpoczeciaMin time.Time
	if len(query["data_roczpoczecia_min"]) > 0 {
		dataRozpoczeciaMin, _ = time.Parse(time.RFC3339, query["data_roczpoczecia_min"][0])
	}
	var dataRozpoczeciaMax time.Time
	if len(query["data_roczpoczecia_max"]) > 0 {
		dataRozpoczeciaMax, _ = time.Parse(time.RFC3339, query["data_roczpoczecia_max"][0])
	}

	var dataZakonczeniaMin time.Time
	if len(query["data_zakonczenia_min"]) > 0 {
		dataZakonczeniaMin, _ = time.Parse(time.RFC3339, query["data_zakonczenia_min"][0])
	}
	var dataZakonczeniaMax time.Time
	if len(query["data_zakonczenia_max"]) > 0 {
		dataZakonczeniaMax, _ = time.Parse(time.RFC3339, query["data_zakonczenia_max"][0])
	}

	imie_pracownika := ""
	if len(query["imie_pracownika"]) > 0 {
		imie_pracownika = query["imie_pracownika"][0]
	}
	nazwisko_pracownika := ""
	if len(query["nazwisko_pracownika"]) > 0 {
		nazwisko_pracownika = query["nazwisko_pracownika"][0]
	}
	marka := ""
	if len(query["marka"]) > 0 {
		marka = query["marka"][0]
	}

	rw.Header().Add("Content-Type", "application/json")

	var kurs []schemas.Kurs
	err := k.db.Model(&kurs).Relation("Trasa").Relation("Kierowca").Relation("Pojazd").Relation("Pojazd.Marka").Relation("Kierowca.Kategorie").
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.WhereOr("data_rozpoczecia >= ?", dataRozpoczeciaMin).WhereOr("TRUE = ?", dataRozpoczeciaMin.IsZero())
			return q, nil
		}).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.WhereOr("data_rozpoczecia <= ?", dataRozpoczeciaMax).WhereOr("TRUE = ?", dataRozpoczeciaMax.IsZero())
			return q, nil
		}).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.WhereOr("data_zakonczenia >= ?", dataZakonczeniaMin).WhereOr("TRUE = ?", dataZakonczeniaMin.IsZero())
			return q, nil
		}).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.WhereOr("data_zakonczenia <= ?", dataZakonczeniaMax).WhereOr("TRUE = ?", dataZakonczeniaMax.IsZero())
			return q, nil
		}).
		Where("kierowca.imie LIKE ?", "%"+imie_pracownika+"%").
		Where("kierowca.nazwisko LIKE ?", "%"+nazwisko_pracownika+"%").
		Where("pojazd__marka.nazwa LIKE ?", "%"+marka+"%").
		Select()
	if err != nil {
		k.l.Error("while handling get all", "path", k.path, "error", err)
		http.Error(rw, "Error getting kurs table", http.StatusInternalServerError)
		return
	}
	for _, jedenKurs := range kurs {
		jedenKurs.Kierowca.Haslo = ""
	}
	err = json.NewEncoder(rw).Encode(kurs)
	if err != nil {
		k.l.Error("err", "", err)
		return
	}
}

func (k *Kursy) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	k.l.Debug("handling get by ID request", "path", k.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	kurs := schemas.Kurs{}
	err = k.db.Model(&kurs).Where("kursy.id = ?", id).Relation("Trasa").Relation("Kierowca").Relation("Pojazd").Relation("Pojazd.Marka").Relation("Kierowca.Kategorie").Select()
	if err != nil {
		k.l.Error("while handling get by ID", "path", k.path, "error", err)
		http.Error(rw, "Error getting kurs table", http.StatusInternalServerError)
		return
	}

	kurs.Kierowca.Haslo = ""

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(kurs)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}

func (k *Kursy) getByDriverID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := r.URL.Query()
	var dataRozpoczeciaMin time.Time
	if len(query["data_roczpoczecia_min"]) > 0 {
		dataRozpoczeciaMin, _ = time.Parse(time.RFC3339, query["data_roczpoczecia_min"][0])
	}
	var dataRozpoczeciaMax time.Time
	if len(query["data_roczpoczecia_max"]) > 0 {
		dataRozpoczeciaMax, _ = time.Parse(time.RFC3339, query["data_roczpoczecia_max"][0])
	}

	var dataZakonczeniaMin time.Time
	if len(query["data_zakonczenia_min"]) > 0 {
		dataZakonczeniaMin, _ = time.Parse(time.RFC3339, query["data_zakonczenia_min"][0])
	}
	var dataZakonczeniaMax time.Time
	if len(query["data_zakonczenia_max"]) > 0 {
		dataZakonczeniaMax, _ = time.Parse(time.RFC3339, query["data_zakonczenia_max"][0])
	}

	imie_pracownika := ""
	if len(query["imie_pracownika"]) > 0 {
		imie_pracownika = query["imie_pracownika"][0]
	}
	nazwisko_pracownika := ""
	if len(query["nazwisko_pracownika"]) > 0 {
		nazwisko_pracownika = query["nazwisko_pracownika"][0]
	}
	marka := ""
	if len(query["marka"]) > 0 {
		marka = query["marka"][0]
	}

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	k.l.Debug("handling get all request", "path", k.path)

	rw.Header().Add("Content-Type", "application/json")

	var kurs []schemas.Kurs
	err = k.db.Model(&kurs).Relation("Trasa").Relation("Kierowca").Relation("Pojazd").Relation("Pojazd.Marka").Relation("Kierowca.Kategorie").
		Where("kursy.kierowca_id = ?", id).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.WhereOr("data_rozpoczecia >= ?", dataRozpoczeciaMin).WhereOr("TRUE = ?", dataRozpoczeciaMin.IsZero())
			return q, nil
		}).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.WhereOr("data_rozpoczecia <= ?", dataRozpoczeciaMax).WhereOr("TRUE = ?", dataRozpoczeciaMax.IsZero())
			return q, nil
		}).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.WhereOr("data_zakonczenia >= ?", dataZakonczeniaMin).WhereOr("TRUE = ?", dataZakonczeniaMin.IsZero())
			return q, nil
		}).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.WhereOr("data_zakonczenia <= ?", dataZakonczeniaMax).WhereOr("TRUE = ?", dataZakonczeniaMax.IsZero())
			return q, nil
		}).
		Where("kierowca.imie LIKE ?", "%"+imie_pracownika+"%").
		Where("kierowca.nazwisko LIKE ?", "%"+nazwisko_pracownika+"%").
		Where("pojazd__marka.nazwa LIKE ?", "%"+marka+"%").Select()
	if err != nil {
		k.l.Error("while handling get all", "path", k.path, "error", err)
		http.Error(rw, "Error getting kurs table", http.StatusInternalServerError)
		return
	}
	for _, jedenKurs := range kurs {
		jedenKurs.Kierowca.Haslo = ""
	}
	err = json.NewEncoder(rw).Encode(kurs)
	if err != nil {
		k.l.Error("err", "", err)
		return
	}
}
