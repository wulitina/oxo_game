package db

/*
import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

// todo the logic of connecting mysql to be done.....
func init() {
	dataSourceName := "root:root@tcp(localhost:3306)/oxo_game?parseTime=true"

	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Check database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Connected to database!")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
*/
