package backend

import (
	"database/sql"
	"log"
)

// Opening the database
func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/social_network.db?parseTime=true")
	if err != nil {
		log.Print(err)
		return nil
	}

	return db
}
