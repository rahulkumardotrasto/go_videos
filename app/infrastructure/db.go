package infrastructure

import (
	"database/sql"
	"log"
)

//Database  Infrastructure layer to incorporate all the database function
type Database struct{}

//GetPostgre ...
func (database *Database) GetPostgre() *sql.DB {
	connStr := "user=postgre dbname=videos sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
