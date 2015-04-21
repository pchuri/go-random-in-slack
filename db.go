package main

import (
	"database/sql"
	"os"

	"github.com/lib/pq"
)

// table scheme
// CREATE TABLE random_logs(id SERIAL PRIMARY KEY, username VARCHAR(40) NOT NULL, random INTEGER NOT NULL)
// CREATE INDEX random_logs_username_idx ON random_logs(username)

func ConnectDB() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	return sql.Open("postgres", connection)
}

func InsertRandomToDB(db *sql.DB, username string, randValue int) error {
	_, err := db.Exec(
		"INSERT INTO random_logs (username, random) VALUES ('?', ?)", username, randValue)

	return err
}
