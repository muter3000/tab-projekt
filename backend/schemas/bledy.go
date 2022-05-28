package schemas

type Blad struct {
	tableName struct{} `pg:"bledy,alias:bledy"`
	Id        int32    `pg:"id,pk" json:"id"`
	Tytul     string   `pg:"tytul" json:"tytul"`
	Opis      string   `pg:"opis" json:"opis"`
}
