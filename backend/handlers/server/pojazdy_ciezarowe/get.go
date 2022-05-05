package pojazdy_ciezarowe

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (pc *PojazdyCiezarowe) getAll(rw http.ResponseWriter, _ *http.Request) {
	pc.l.Debug("handling get all request", "path", pc.path)

	rw.Header().Add("Content-Type", "application/json")

	var pojazd_ciezarowy []schemas.PojazdCiezarowy
	err := pc.db.Model(&pojazd_ciezarowy).Select()
	if err != nil {
		pc.l.Error("while handling get all", "path", pc.path, "error", err)
		http.Error(rw, "Error getting pojazd_ciezarowy table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(pojazd_ciezarowy)
	if err != nil {
		pc.l.Error("err", "", err)
		return
	}
}

func (pc *PojazdyCiezarowe) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	pc.l.Debug("handling get by ID request", "path", pc.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	pojazd_ciezarowy := schemas.PojazdCiezarowy{}
	err = pc.db.Model(&pojazd_ciezarowy).Where("id = ?", id).Select()
	if err != nil {
		pc.l.Error("while handling get by ID", "path", pc.path, "error", err)
		http.Error(rw, "Error getting pojazd_ciezarowy table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(pojazd_ciezarowy)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
