package db


import (
	"database/sql"
	_ "github.com/lib/pq"
)


type Database struct{
	Db *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func(db *Database) StartDb() error {

	sqldb, err := sql.Open("postgres", "postgres://postgres:Dsquare142@localhost/restaurant?sslmode=disable")
	
	if err != nil {
		return err
	}
	db.Db = sqldb
	return nil
}


