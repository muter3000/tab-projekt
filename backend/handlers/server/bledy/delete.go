package bledy

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tab-projekt-backend/schemas"
)

func (b *Bledy) deleteByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	b.l.Debug("handling delete by ID request", "path", b.path, "id", id)

	rw.Header().Add("Content-Type", "application/json")

	blad := schemas.Blad{}
	deleted, err := b.db.Model(&blad).Where("id = ?", id).Delete()
	if err != nil {
		b.l.Error("while handling delete by ID", "path", b.path, "error", err)
		http.Error(rw, "Error getting blad table", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(deleted.RowsAffected())
	if err != nil {
		http.Error(rw, "Error encoding to json", http.StatusInternalServerError)
		return
	}
}
