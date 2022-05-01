package schemas

import "time"

type Trasa struct {
	tableName             struct{} `pg:"trasy"`
	Id                    int32    `pg:"id,pk" json:"id"`
	MiejscowoscPoczatkowa string   `pg:"miejscowosc_poczatkowa" json:"miejscowosc_poczatkowa"`
	MiejscowoscKoncowa    string   `pg:"miejscowosc_koncowa" json:"miejscowosc_koncowa"`
}

type Kurs struct {
	tableName       struct{}  `pg:"kursy"`
	Id              int32     `pg:"id,pk" json:"id"`
	DataRozpoczecia time.Time `pg:"data_rozpoczecia" json:"data_rozpoczecia"`
	DataZakonczenia time.Time `pg:"data_zakonczenia" json:"data_zakonczenia"`
	CzasRozpoczecia time.Time `pg:"czas_rozpoczecia" json:"czas_rozpoczecia"`
	Ladunek         float32   `pg:"ladunek" json:"ladunek"`
	TrasaID         int32     `pg:",notnull,on_delete:RESTRICT"`
	Trasa           *Trasa    `pg:"rel:has-one"`
	KierowcaID      int32     `pg:",notnull,on_delete:RESTRICT"`
	Kierowca        *Kierowca `pg:"rel:has-one"`
	PojazdID        int32     `pg:",notnull,on_delete:RESTRICT"`
	Pojazd          *Pojazd   `pg:"rel:has-one"`
}
