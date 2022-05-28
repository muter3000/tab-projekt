package main

import (
	"fmt"

	"github.com/go-pg/pg/v10/orm"
	"github.com/hashicorp/go-hclog"
	"github.com/tab-projekt-backend/database/psql"
	"github.com/tab-projekt-backend/schemas"
)

func main() {
	log := hclog.Default()
	db, err := psql.GetDB()
	if err != nil {
		log.Error("connecting to db", "err", err)
	}
	defer db.Close()

	orm.RegisterTable((*schemas.KategoriaKierowcy)(nil))

	models := []interface{}{
		(*schemas.Pracownik)(nil),
		(*schemas.StanowiskoAdministracyjne)(nil),
		(*schemas.KategoriaPrawaJazdy)(nil),
		(*schemas.KategoriaKierowcy)(nil),
		(*schemas.Kierowca)(nil),
		(*schemas.Administrator)(nil),
		(*schemas.Marka)(nil),
		(*schemas.Pojazd)(nil),
		(*schemas.PojazdCiezarowy)(nil),
		(*schemas.Trasa)(nil),
		(*schemas.Kurs)(nil),
		(*schemas.Blad)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true, FKConstraints: true})
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("Database created\n")
}
