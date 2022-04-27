package schemas

type Pojazd struct {
	tableName        struct{} `pg:"pojazdy"`
	Id               int32    `pg:"id,pk" json:"id"`
	NumerSilnika     int32    `pg:"numer_silnika" json:"numer_silnika"`
	PojemnoscSilnika int32    `pg:"pojemnosc_silnika" json:"pojemnosc_silnika"`
	MarkaID          int32    `json:"marka_id"`
	Marka            *Marka   `pg:"rel:has-one"`
}

type PojazdCiezarowy struct {
	Pojazd    `pg:",inherit"`
	tableName struct{} `pg:"pojazdy_ciezarowe"`
	Ladownosc float32  `pg:"ladownosc" json:"ladownosc"`
}

type Marka struct {
	tableName struct{} `pg:"marki"`
	Id        int32    `pg:"id,pk" json:"id"`
	Nazwa     string   `pg:"nazwa" json:"nazwa"`
}
