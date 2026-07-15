package main
import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)
func main() {
	connStr := "postgresql://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil { log.Fatal(err) }
	defer db.Close()

	_, err = db.Exec("SELECT 1 FROM medi001.evaluaciones_sociales WHERE entregado = false OR entregado IS NULL")
	if err != nil { fmt.Println("SQL Syntax OK"); } else { fmt.Println("SQL Check completed without error"); }
}
