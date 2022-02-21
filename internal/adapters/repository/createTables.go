package repository

import "database/sql"

func CreateTables(db *sql.DB) error {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			userName TEXT UNIQUE,
			eMail TEXT UNIQUE,
			password TEXT
		);
	`)

	if err != nil {
		return err
	}
	_, err = db.Exec(` 
		CREATE TABLE IF NOT EXISTS Cookie(
			value TEXT,
			userId INTEGER UNIQUE,
			expires DATETIME not null,
			FOREIGN KEY (userId)
				REFERENCES Users (id)
		)
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Posts(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		body TEXT,
		userId INTEGER,
		FOREIGN KEY (userId)
			REFERENCES Cookie (userId)
	)`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Comments(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		body TEXT,
		userId INTEGER,
		postId INTEGER,
		FOREIGN KEY (postId)
			REFERENCES Posts (id)
	)`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS DisLikes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		like INTEGER,
		dislike INTEGER,
		userId INTEGER,
		postIdInComm INTEGER,
		FOREIGN KEY (postIdInComm)
			REFERENCES Comments (postId)
	)`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS CommDisLikes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		like INTEGER,
		dislike INTEGER,
		userId INTEGER,
		commId INTEGER,
		FOREIGN KEY (commId)
			REFERENCES Comments (postId)
	)`)
	if err != nil {
		return err
	}

	return nil
}
