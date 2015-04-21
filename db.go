package main

import (
	"database/sql"
	"os"

	"github.com/lib/pq"
)

// table scheme
// CREATE TABLE random_logs(id SERIAL PRIMARY KEY, username VARCHAR(40) NOT NULL, randvalue INTEGER NOT NULL)
// CREATE INDEX random_logs_username_idx ON random_logs(username)

func ConnectDB() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	return sql.Open("postgres", connection)
}

func InsertRandomToDB(db *sql.DB, username string, randValue int) error {
	stmt, err := db.Prepare("INSERT INTO random_logs (\"username\", \"randvalue\") VALUES ($1,$2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, randValue)

	return err
}

func SelectAvgRandomFromDB(db *sql.DB, username string) (int, error) {
	stmt, err := db.Prepare("SELECT avg(randvalue) from random_logs where username = $1")
	if err != nil {
		return 0.0, err
	}

	row := stmt.QueryRow(username)

	var avgRand int
	err = row.Scan(&avgRand)

	return avgRand, err
}
