package postgres

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"os"
)

func GetDB() (*pg.DB, error) {
	dbHost := os.Getenv("PSQL_HOST")
	dbPort := os.Getenv("PSQL_PORT")
	dbUser := os.Getenv("PSQL_USER")
	dbPassword := os.Getenv("PSQL_PASSWORD")
	dbName := os.Getenv("PSQL_DB")

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})

	ctx := context.Background()

	var version string
	_, err := db.QueryOneContext(ctx, pg.Scan(&version), "SELECT version()")
	if err != nil {
		return nil, err
	}

	return db, nil
}
