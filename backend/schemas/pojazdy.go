package schemas

type Pojazd struct {
	tableName          struct{} `pg:"pojazdy, alias:pojazdy"`
	Id                 int32    `pg:"id, pk" json:"id"`
	Vin                string   `pg:"vin" json:"vin"`
	PojemnoscSilnika   int32    `pg:"pojemnosc_silnika" json:"pojemnosc_silnika"`
	MarkaID            int32    `pg:", notnull,on_delete:RESTRICT" json:"marka_id"`
	NumerRejestracyjny string   `pg:"numer_rejestracyjny, unique" json:"numer_rejestracyjny"`
	Model              string   `pg:"model" json:"model"`
	Marka              *Marka   `pg:"rel:has-one"`
}

type PojazdCiezarowy struct {
	Pojazd    `pg:", inherit"`
	tableName struct{} `pg:"pojazdy_ciezarowe,alias:pojazdy_ciezarowe"`
	Ladownosc float32  `pg:"ladownosc" json:"ladownosc"`
}

type Marka struct {
	tableName struct{} `pg:"marki"`
	Id        int32    `pg:"id, pk" json:"id"`
	Nazwa     string   `pg:"nazwa" json:"nazwa"`
}
