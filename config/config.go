package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenConn() *sqlx.DB {
	db, err := sqlx.Open("postgres", "user=postgres password=althea dbname=AlthCart sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(0)
	db.SetConnMaxLifetime(0)

	return db
}
