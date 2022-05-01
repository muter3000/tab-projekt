package kategoria_prawa_jazdy

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (kpj *Kategoria_prawa_jazdy) getAll(rw http.ResponseWriter, _ *http.Request) {
	kpj.l.Debug("handling get all request", "path", kpj.path)

	rw.Header().Add("Content-Type", "application/json")

	var kategorie []schemas.KategoriaPrawaJazdy
	err := kpj.db.Model(&kategorie).Select()
	if err != nil {
		kpj.l.Error("while handling get all", "path", kpj.path, "error", err)
		http.Error(rw, "Error getting kategorie table", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(kategorie)
	if err != nil {
		kpj.l.Error("err", "", err)
		return
	}
}

func (kpj *Kategoria_prawa_jazdy) getByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	kpj.l.Debug("handling get by ID request", "path", kpj.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	kategoria := schemas.KategoriaPrawaJazdy{Id: int32(id)}
	err = kpj.db.Model(&kategoria).Select()
	if err != nil {
		kpj.l.Error("while handling get by ID", "path", kpj.path, "error", err)
		http.Error(rw, "Error getting kategorie table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(kategoria)
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
