package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	// connect to database
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}

	DB = db
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time TEXT NOT NULL,
		user_id INTEGER NOT NULL
	);`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}
