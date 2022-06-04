package main

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/hashicorp/go-hclog"
	"github.com/tab-projekt-backend/database/psql"
	"github.com/tab-projekt-backend/schemas"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	value, _ := q.FormattedQuery()
	fmt.Println(string(value))
	return nil
}

func main() {
	log := hclog.Default()
	db, err := psql.GetDB()
	if err != nil {
		log.Error("connecting to db", "err", err)
	}
	defer db.Close()

	db.AddQueryHook(dbLogger{})

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
			fmt.Print(err)
			panic(err)
		}
	}
	fmt.Printf("Database created\n")
}
