// begin open OMIT
package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // driver package for PostgreSQL
)

func main() {
	db, err := sql.Open("postgres", "user=test dbname=test sslmode=disable") // HL
	if err != nil {
		log.Fatal(err)
	}

	// end open OMIT
	// begin query OMIT
	var (
		id   int
		name string
	)

	rows, err := db.Query(`select id, name from users where id = ?`, 1) // HL
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // HL
	for rows.Next() {  // HL
		err := rows.Scan(&id, &name) // HL
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err() // HL
	if err != nil {
		log.Fatal(err)
	}
	// end query OMIT
}
