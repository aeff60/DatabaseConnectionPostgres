package main
import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)
const (
  host     = "localhost"
  port     = 5433
  user     = "postgres"
  password = "mflv[DB2023"
  dbname   = "CompanyDB"
)
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	query := `SELECT * FROM "CompanyDB"."public"."employee"`
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var employeeid int
		var fistname string
		var lastname string
		err = rows.Scan(&employeeid, &fistname, &lastname)
		if err != nil {
			panic(err)
		}
		fmt.Println(employeeid, fistname, lastname)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}