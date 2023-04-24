package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	Postgre   *sql.DB
	username = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	schema   = os.Getenv("DB_NAME")
	host = os.Getenv("DB_HOST")
)
func init() {
	connInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host,	username, password, schema)
	Postgre, err := sql.Open("pgx", connInfo)
	if err != nil {
		log.Fatalf("Error %v while connecting to postgre db", err)
	}
	if err = Postgre.Ping(); err != nil {
		log.Fatalf("Error %v while pinging postgre db", err)
	} 
	log.Println("Postgre database ready to accept connections")
}