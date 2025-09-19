package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

const SCHEMA = `
PRAGMA journal_mode = WAL;
PRAGMA busy_timeout = 5000;

CREATE TABLE IF NOT EXISTS ingredients (
	ingredient_id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	cost REAL NOT NULL
);
`

func OpenDatabase(dbPath string) error {
	var err error
	db, err = sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	db.MustExec(SCHEMA)
	return nil
}

func CloseDatabase() {
	if db != nil {
		db.Close()
	}
}
