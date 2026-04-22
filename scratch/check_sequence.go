package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var lastId int
	err = db.QueryRow("SELECT MAX(id) FROM medi001.doctores").Scan(&lastId)
	if err != nil {
		log.Fatal(err)
	}

	var nextSeq int
	err = db.QueryRow("SELECT nextval('medi001.doctores_id_seq')").Scan(&nextSeq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Max ID: %d, Next Seq: %d\n", lastId, nextSeq)
}
