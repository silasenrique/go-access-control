package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic("não foi possível pingar o banco")
	}

	return db
}
