package kursy

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (k *Kursy) getAll(rw http.ResponseWriter, _ *http.Request) {
	k.l.Debug("handling get all request", "path", k.path)

	rw.Header().Add("Content-Type", "application/json")

	var kurs []schemas.Kurs
	err := k.db.Model(&kurs).Relation("Trasa").Relation("Kierowca").Relation("Pojazd").Relation("Pojazd.Marka").Relation("Kierowca.Kategorie").Select()
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

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	k.l.Debug("handling get all request", "path", k.path)

	rw.Header().Add("Content-Type", "application/json")

	var kurs []schemas.Kurs
	err = k.db.Model(&kurs).Where("kursy.kierowca_id = ?", id).Relation("Trasa").Relation("Kierowca").Relation("Pojazd").Select()
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
