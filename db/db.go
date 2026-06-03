package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DBcon *sqlx.DB

func Connect() error {

	user := "adonia"
	password := "adonia"
	host := "localhost"
	port := "5432"
	dbname := "internship"
	sslmode := "disable"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	DBcon = db
	return nil
}
