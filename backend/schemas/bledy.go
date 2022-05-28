package schemas

import "time"

type Blad struct {
	tableName      struct{}  `pg:"bledy,alias:bledy"`
	Id             int32     `pg:"id,pk" json:"id"`
	Tytul          string    `pg:"tytul" json:"tytul"`
	Opis           string    `pg:"opis" json:"opis"`
	Kategoria      string    `pg:"kategoria" json:"kategoria"`
	DataUtworzenia time.Time `pg:"data_utworzenia, default:now()" json:"data_utworzenia"`
}
