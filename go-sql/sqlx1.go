// begin open OMIT
package main

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // driver package for PostgreSQL
)

func main() {
	db, err := sqlx.Open("postgres", "user=test dbname=test sslmode=disable") // HL
	if err != nil {
		log.Fatal(err)
	}

	// end open OMIT
	// begin query OMIT
	var user struct {
		ID   int
		Name string
	}

	err = db.Get(&user, `select id, name from users where id = ?`, 1) // HL
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user.ID, user.Name)
	// end query OMIT
}
