package repository

import "database/sql"

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "dataBase.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
