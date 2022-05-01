package schemas

import (
	"context"
	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
)

type Pracownik struct {
	tableName struct{} `pg:"pracownicy"`
	Id        int32    `pg:"id,pk" json:"id"`
	Pesel     string   `pg:"pesel" json:"pesel"`
	Imie      string   `pg:"imie" json:"imie"`
	Nazwisko  string   `pg:"nazwisko" json:"nazwisko"`
	Login     string   `pg:"login" json:"login"`
	Haslo     string   `pg:"haslo" json:"haslo,omitempty"`
}

//Compile time check
var _ pg.AfterScanHook = (*Pracownik)(nil)

//

func (p *Pracownik) AfterScan(context.Context) error {
	p.Haslo = ""
	return nil
}

var _ pg.BeforeInsertHook = (*Pracownik)(nil)

func (p *Pracownik) BeforeInsert(ctx context.Context) (context.Context, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(p.Haslo), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	p.Haslo = string(password)
	return ctx, nil
}

type Kierowca struct {
	Pracownik  `pg:",inherit"`
	tableName  struct{} `pg:"kierowcy"`
	KierowcaID int32    `pg:"kierowca_id,pk" json:"kierowca_id"`
}

type Administrator struct {
	Pracownik                   `pg:",inherit"`
	tableName                   struct{}                   `pg:"administratorzy"`
	AdministracjaID             int32                      `pg:"administracja_id,pk" json:"administracja_id"`
	StanowiskoAdministracyjneID int32                      `json:"stanowisko_administracyjne_id"`
	StanowiskoAdministracyjne   *StanowiskoAdministracyjne `pg:"rel:has-one"`
}
type StanowiskoAdministracyjne struct {
	tableName struct{} `pg:"stanowiska_administracyjne"`
	Id        int32    `pg:"id,pk" json:"id"`
	Typ       string   `pg:"typ" json:"typ"`
}

type KategoriaPrawaJazdy struct {
	tableName struct{} `pg:"kategorie_prawa_jazdy"`
	Id        int32    `pg:"id,pk" json:"id"`
	Kategoria string   `pg:"kategoria" json:"kategoria"`
}

type KategoriaKierowcy struct {
	tableName             struct{}             `pg:"kategoria_kierowcy"`
	KierowcaID            int32                `json:"kierowca_id"`
	Kierowca              *Kierowca            `pg:"rel:has-one"`
	KategoriaPrawaJazdyID int32                `json:"kategoria_prawa_jazdy_id"`
	KategoriaPrawaJazdy   *KategoriaPrawaJazdy `pg:"rel:has-one"`
}
