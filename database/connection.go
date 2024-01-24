package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "testdb"
)

func Connect() *sql.DB {
	// connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
	db, err := sql.Open("postgres", connStr)
	CheckError(err)

	fmt.Println("Connected!")
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
