package kursy

import (
	"encoding/json"
	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/helpers"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (k *Kursy) updateExisting(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	rw.Header().Add("Content-Type", "application/json")
	k.l.Debug("handling post request", "path", k.path)
	kurs := schemas.Kurs{}
	err = json.NewDecoder(r.Body).Decode(&kurs)
	kurs.Id = int32(id)
	if err != nil || kurs.CzasRozpoczecia == nil || kurs.CzasZakonczenia == nil || kurs.CzasPrzejazdu == nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	_, err = k.db.Model(&kurs).Column("czas_rozpoczecia").Column("czas_zakonczenia").Column("czas_przejazdu").WherePK().Returning("*", &kurs).Update()
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(kurs)
	if err != nil {
		http.Error(rw, "Error marshaling new kurs", http.StatusBadRequest)
		return
	}

}

func (k *Kursy) updateMyExisting(rw http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetAuthAndIDFromSession(r, redis.Kierowca)
	if err != nil {
		http.Error(rw, "No session currently active", http.StatusUnauthorized)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	k.l.Debug("handling post request", "path", k.path)
	kurs := schemas.Kurs{}
	err = json.NewDecoder(r.Body).Decode(&kurs)
	kurs.Id = int32(id)
	if err != nil || kurs.CzasRozpoczecia == nil || kurs.CzasZakonczenia == nil || kurs.CzasPrzejazdu == nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Wrong parameters passed", http.StatusBadRequest)
		return
	}
	_, err = k.db.Model(&kurs).Column("czas_rozpoczecia").Column("czas_zakonczenia").Column("czas_przejazdu").WherePK().Returning("*", &kurs).Update()
	if err != nil {
		k.l.Error("marshaling", "err", err)
		http.Error(rw, "Inserting into database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(kurs)
	if err != nil {
		http.Error(rw, "Error marshaling new kurs", http.StatusBadRequest)
		return
	}

}
