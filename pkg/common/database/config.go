package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	POSTGRES_DB_URL := "postgres://postgres:admin@127.0.0.1:5432/it_parser?sslmode=disable"

	db, err := sql.Open("postgres", POSTGRES_DB_URL)
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
