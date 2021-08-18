package lexicon

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

var Db *DB

func ConnectPsql() {
	log.Println("Instanciating the database...")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"))

	conn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Failed to open DB connection: ", err)
	}

	Db = &DB{}
	Db.conn = conn
	log.Println("Database connected.")

}

func (t *DB) queryWord(model string, current string) *sql.Rows {
	query := fmt.Sprintf(model, current)
	r, err := t.conn.Query(query)
	if err != nil {
		panic(err)
	}
	return r
}
