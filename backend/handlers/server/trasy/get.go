package trasy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (t *Trasy) getAll(rw http.ResponseWriter, _ *http.Request) {
	t.l.Debug("handling get all request", "path", t.path)

	rw.Header().Add("Content-Type", "application/json")

	var trasa []schemas.Trasa
	err := t.db.Model(&trasa).Select()
	if err != nil {
		t.l.Error("while handling get all", "path", t.path, "error", err)
		http.Error(rw, "Error getting trasa table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(trasa)
	if err != nil {
		t.l.Error("err", "", err)
		return
	}
}

func (t *Trasy) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	t.l.Debug("handling get by ID request", "path", t.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	trasa := schemas.Trasa{}
	err = t.db.Model(&trasa).Where("id = ?", id).Select()
	fmt.Print(trasa)
	if err != nil {
		t.l.Error("while handling get by ID", "path", t.path, "error", err)
		http.Error(rw, "Error getting trasa table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(trasa)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
