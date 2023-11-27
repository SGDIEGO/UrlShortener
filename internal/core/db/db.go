package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const driver = "sqlite3"
const path_db = "../internal/database/db_data"

func LoadDb() (db *sql.DB, e error) {
	db_, err := sql.Open(driver, path_db)
	if err != nil {
		return nil, err
	}

	err = db_.Ping()
	if err != nil {
		return nil, err
	}

	//	err = db_.Close()
	//	if err != nil {
	//		return nil, err
	//	}

	return db_, nil
}
