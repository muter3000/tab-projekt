package schemas

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
)

type Pracownik struct {
	tableName struct{} `pg:"pracownicy"`
	Id        int32    `pg:"id,pk" json:"id"`
	Pesel     string   `pg:"pesel,unique" json:"pesel"`
	Imie      string   `pg:"imie" json:"imie"`
	Nazwisko  string   `pg:"nazwisko" json:"nazwisko"`
	Login     string   `pg:"login,unique" json:"login"`
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

var _ pg.AfterInsertHook = (*Pracownik)(nil)

func (p *Pracownik) AfterInsert(context.Context) error {
	p.Haslo = ""
	return nil
}

type Kierowca struct {
	Pracownik  `pg:",inherit"`
	tableName  struct{} `pg:"kierowcy"`
	KierowcaID int32    `pg:"kierowca_id,pk" json:"kierowca_id"`

	Kategorie []KategoriaPrawaJazdy `pg:"many2many:kategoria_kierowcy,fk:kierowca_" json:"kategorie"`
}

var _ pg.BeforeInsertHook = (*Kierowca)(nil)

func (k *Kierowca) BeforeInsert(ctx context.Context) (context.Context, error) {
	valid, err := k.validate()
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, errors.New("user input from post request invalid")
	}

	ctx2, err := k.Pracownik.BeforeInsert(ctx)
	if err != nil {
		return nil, err
	}
	return ctx2, nil
}

type Administrator struct {
	Pracownik                   `pg:",inherit"`
	tableName                   struct{}                   `pg:"administratorzy"`
	AdministracjaID             int32                      `pg:"administracja_id,pk" json:"administracja_id"`
	StanowiskoAdministracyjneID int32                      `pg:",notnull,on_delete:RESTRICT" json:"stanowisko_administracyjne_id"`
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
	tableName             struct{} `pg:"kategoria_kierowcy"`
	KierowcaId            int32    `pg:",notnull,on_delete:RESTRICT,fk:kierowca_id,join_fk:kierowca_id" json:"kierowca_id"`
	KategoriaPrawaJazdyId int32    `pg:",notnull,on_delete:RESTRICT,fk:id,join_fk:id" json:"kategoria_prawa_jazdy_id"`
}
